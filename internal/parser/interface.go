package parser

import (
	"steplife-universal-importer-gui/internal/model"
	xif "steplife-universal-importer-gui/internal/utils/if"
	"steplife-universal-importer-gui/internal/utils/logx"
	"steplife-universal-importer-gui/internal/utils/pointcalc"
)

type FileAdaptor interface {
	//
	// Parse
	//  @Description: 		文件解析
	//  @param content
	//  @return []float64	返回经纬度坐标
	//  @return error
	//
	Parse(content []byte) ([]model.Point, error)

	//
	// Convert2StepLife
	//  @Description: 			将经纬度坐标转换成一生足迹数据结构
	//  @param config 			路径转换配置信息
	//  @param points
	//  @return *model.StepLife
	//  @return error
	Convert2StepLife(config model.Config, points []model.Point) (*model.StepLife, error)
}

type BaseAdaptor struct{}

func (this *BaseAdaptor) Parse(content []byte) ([]model.Point, error) {
	panic("implement me")
}

func (this *BaseAdaptor) Convert2StepLife(config model.Config, points []model.Point) (*model.StepLife, error) {
	previousPoint := model.Point{}

	sl := model.NewStepLife()
	logx.Info("处理经纬度")
	
	// 检查时间设置
	// 优先级：结束时间 > 用户指定的时间间隔 > 统一为开始时间
	useEndTime := config.PathEndTimestamp > 0 && config.PathStartTimestamp > 0
	useTimeInterval := config.TimeInterval != 0 && config.PathStartTimestamp > 0
	
	// 计算时间间隔（如果设置了结束时间）
	var timeInterval int64 = 1
	if useEndTime {
		// 需要先计算总点数
		totalPoints := int64(len(points))
		if config.EnableInsertPointStrategy == 1 {
			totalPoints = 1
			for i := 1; i < len(points); i++ {
				interpolatedPoints := pointcalc.Calculate(points[i-1], points[i], config.InsertPointDistance)
				totalPoints += int64(len(interpolatedPoints))
			}
		}
		if totalPoints > 1 {
			totalDuration := config.PathEndTimestamp - config.PathStartTimestamp
			timeInterval = totalDuration / (totalPoints - 1)
			if timeInterval < 1 {
				timeInterval = 1
			}
		}
	} else if useTimeInterval {
		timeInterval = config.TimeInterval
	}
	
	pointIndex := int64(0)
	
	for i, point := range points {

		// 第0个坐标或者不需要插入值，不需要计算中间点，直接写入
		if i == 0 || config.EnableInsertPointStrategy == 0 {
			row := model.NewRow()
			if useEndTime {
				// 如果设置了结束时间，使用计算出的时间间隔
				currentTimestamp := config.PathStartTimestamp + pointIndex*timeInterval
				if i == len(points)-1 {
					currentTimestamp = config.PathEndTimestamp
				}
				point.DataTime = xif.Int64(point.DataTime == 0, currentTimestamp, point.DataTime)
			} else if useTimeInterval {
				// 如果设置了时间间隔，按照间隔递增
				currentTimestamp := config.PathStartTimestamp + pointIndex*timeInterval
				point.DataTime = xif.Int64(point.DataTime == 0, currentTimestamp, point.DataTime)
			} else {
				// 如果都没有设置，所有时间统一为开始时间
				point.DataTime = xif.Int64(point.DataTime == 0, config.PathStartTimestamp, point.DataTime)
			}
			row.Point = point
			sl.AddCSVRow(*row)
			pointIndex++
		} else {
			interpolatedPoints := pointcalc.Calculate(previousPoint, point, config.InsertPointDistance)
			for j, interpolatedPoint := range interpolatedPoints {
				row := model.NewRow()
				row.Point = interpolatedPoint
				if useEndTime {
					// 如果设置了结束时间，使用计算出的时间间隔
					currentTimestamp := config.PathStartTimestamp + pointIndex*timeInterval
					if i == len(points)-1 && j == len(interpolatedPoints)-1 {
						currentTimestamp = config.PathEndTimestamp
					}
					row.DataTime = xif.Int64(
						interpolatedPoint.DataTime == 0,
						currentTimestamp,
						interpolatedPoint.DataTime,
					)
				} else if useTimeInterval {
					// 如果设置了时间间隔，按照间隔递增
					currentTimestamp := config.PathStartTimestamp + pointIndex*timeInterval
					row.DataTime = xif.Int64(
						interpolatedPoint.DataTime == 0,
						currentTimestamp,
						interpolatedPoint.DataTime,
					)
				} else {
					// 如果都没有设置，所有时间统一为开始时间
					row.DataTime = xif.Int64(
						interpolatedPoint.DataTime == 0,
						config.PathStartTimestamp,
						interpolatedPoint.DataTime,
					)
				}
				sl.AddCSVRow(*row)
				pointIndex++
			}
		}
		previousPoint = point
	}
	logx.InfoF("处理经纬度完成，原始坐标%d个，插点后坐标%d个", len(points), len(sl.CSVData))
	return sl, nil
}

func CreateAdaptor(parserType string) FileAdaptor {
	switch parserType {
	case ".kml":
		return NewKMLAdaptor()
	case ".ovjsn":
		return NewOvjsnAdaptor()
	case ".gpx":
		return NewGpxAdaptor()
	default:
		return nil
	}
}
