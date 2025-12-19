<p align="center">
	<strong>用最快的方式，将三方数据轨迹导入到「一生足迹」中</strong>
</p> 
<h1 align="center">一生足迹数据导入器</h1>


## 1. 项目简介

**一生足迹数据导入器** 是一个便捷的数据转换工具，支持将第三方轨迹数据（如奥维互动地图、KML 等）快速转换为一生足迹 App 所需的 CSV 格式，帮助你更完整地记录旅途与人生的点点滴滴。

> 演示视频
> 
> https://github.com/kearney/steplife-universal-importer/blob/main/static/vedio/demo.mp4
> 
## 2. 前言

### 2.1 一生足迹 · 记录一生轨迹

> 你好，
我是足迹的开发者。<br>
2017 年的清明节（17.4.4），我和同事去了一次沙漠，徒步 12 小时，25.7 公里，我们穿越了中国第七大沙漠库布其。
这次旅行很累，也让我记忆深刻。<br>
在回北京的路上有人在群里问，我们是从哪儿走到哪儿的？<br>
可遗憾的是，没有一个人能准确说出。<br>
突然我就感觉挺悲哀的。<br>
也许我的一生也会这样，不管经历多少，遭遇了多少，到最后，都会被淡忘。没有人记得我们的过去。就像我们这次的旅行一样，很快就会被忘记。<br>
所以，我写了这个可以记录一生轨迹的 APP。<br>
—— By 足迹开发者

这是一款专注于个人轨迹记录的 App。多次获得 Apple App Store 推荐，数据完全保存在本地，无服务器、无广告、隐私安全可靠。

虽然一生足迹的开发者在耗电控制与定位漂移优化方面做了大量努力，已经显著提升了使用体验，但在某些极端环境下，仍可能出现记录不完整的问题。

你是否也遇到过这样的场景：
- 高铁出行，信号差
- 飞机出行，轨迹中断
- 沙漠深山中，GPS 不稳定
- 手机电量不足
- ......

为了解决这些特殊场景下的记录缺失问题，「数据导入器」应运而生。

## 3. 核心特性

- ✅ 支持 **奥维互动地图** 的 Omap JSON 格式导入
- ✅ 支持 **标准 KML 文件格式** 的轨迹导入
- ✅ 支持 **标准 GPX 文件格式** 的轨迹导入
- ✅ 一键转换为一生足迹所需的 CSV 格式
- ✅ **全新GUI界面**，直观易用，支持实时参数调整
- ✅ **自定义字体**，使用小米MiSans字体提供优质视觉体验
- ✅ **批量文件处理**，支持文件夹内多文件自动处理
- ✅ **智能时间分配**，支持设置开始和结束时间
- ✅ **海拔高度控制**，支持自定义默认海拔
- ✅ **速度计算引擎**，自动计算或手动指定速度
- ✅ 无需联网，安全可靠
- ✅ 轨迹合并功能
- 🛠️ 未来将支持「飞常准」等更多格式

## 4. 高级特性

### 4.1 配置概览

本项目支持一系列配置选项，帮助用户根据需求调整轨迹处理与路径优化的参数。以下是各配置项的详细说明：

```ini
# 是否启用插点策略。0: 不启用；1: 启用。缺省值1
enableInsertPointStrategy = 1

# 插点间距。当两点间距离大于该值时，认为两点之间应该插点。单位：米。缺省值100。最小值30
insertPointDistance = 100

# 路径开始时间。格式：yyyy-MM-dd HH:mm:ss。缺省值为当前系统时间
# pathStartTime = 2018-01-01 00:00:00
```

#### 4.1.1 轨迹精度优化

为了增强轨迹的连贯性和可视效果，本项目在处理第三方轨迹数据时内置了轨迹补点算法。该算法会根据轨迹点之间的间距进行智能插点，确保数据的可视化效果更为平滑。

**功能说明：**

- 当轨迹中相邻两个点之间的距离超过设定的阈值（默认 100 米）时，程序会自动插入补点，使得轨迹在地图上更为平滑。
- 该功能特别适用于高铁、飞机等定位信号间断较长的场景。

**默认情况下，插点策略是启用的**，如果不需要自动补点，可通过调整配置关闭此功能。

🔧 **关闭轨迹补点的示例：**

```ini
# 是否启用插点策略。0: 不启用；1: 启用。缺省值1
enableInsertPointStrategy = 0

# 插点间距。当两点间距离大于该值时，认为两点之间应该插点。单位：米。缺省值100。最小值30
insertPointDistance = 100
```

📌 提示：插值点将采用直线插值方式生成，确保轨迹在地图上的视觉连续性，但不会对原始数据精度产生影响。

#### 4.1.2 时间设置

支持设置轨迹的开始时间和结束时间，用于精确控制轨迹的时间分配。

**功能说明：**

- `pathStartTime`: 轨迹开始时间，格式 `yyyy-MM-dd HH:mm:ss`
- `pathEndTime`: 轨迹结束时间，格式 `yyyy-MM-dd HH:mm:ss`（可选）
- 如果设置了结束时间，系统会自动在开始时间和结束时间之间均匀分配轨迹点的时间戳
- 如果只设置开始时间，则使用系统默认的时间间隔

**示例：**

```ini
# 轨迹开始时间
pathStartTime = 2024-04-21 08:00:00
# 轨迹结束时间（可选，用于时间分配）
pathEndTime = 2024-04-21 18:00:00
```

#### 4.1.3 海拔高度设置

设置轨迹点的默认海拔高度。

**功能说明：**

- `defaultAltitude`: 默认海拔高度，单位米
- 所有轨迹点将使用此海拔高度值

**示例：**

```ini
# 默认海拔高度（米）
defaultAltitude = 316.0
```

#### 4.1.4 速度设置

控制轨迹中每个点的速度信息。

**功能说明：**

- `speedMode`: 速度计算模式
  - `auto`: 自动根据轨迹点间距离估算速度
  - `manual`: 使用手动指定的固定速度
- `manualSpeed`: 手动指定速度值（m/s），仅在manual模式下有效

**示例：**

```ini
# 速度计算模式：auto 或 manual
speedMode = auto
# 手动指定速度（m/s），仅在speedMode=manual时有效
manualSpeed = 1.5
```

#### 4.1.5 批量处理设置

控制是否启用批量文件处理功能。

**功能说明：**

- `enableBatchProcessing`: 是否启用批量处理
  - `1`: 启用，支持处理文件夹内所有支持的文件
  - `0`: 禁用，仅处理单个文件

**示例：**

```ini
# 是否启用批量处理
enableBatchProcessing = 1
```


## 5. 快速开始

### 5.1 下载和构建

#### 方式一：下载预构建版本
访问[Releases 页面](https://github.com/hygao1024/steplife-universal-importer/releases)下载最新版本，然后解压到任意目录。

#### 方式二：从源码构建

**环境要求：**
- Go 1.24.2 或更高版本
- 支持的操作系统：Windows、macOS、Linux

**构建步骤：**

```bash
# 克隆项目
git clone https://github.com/hygao1024/steplife-universal-importer.git
cd steplife-universal-importer

# 下载依赖
go mod tidy

# 构建
go build -o main ./cmd
```

**或者使用构建脚本：**
- Linux/macOS: `./build.sh`
- Windows: `build.bat`

### 5.4 运行测试

项目包含完整的测试套件，位于 `tests/` 目录：

```bash
# 快速测试（推荐）
./tests/quick_test.sh

# 单独测试
./tests/test_cli.sh      # 命令行模式测试
./tests/test_gui.sh      # GUI测试指南
./tests/verify_output.sh # 输出验证
./tests/demo_gui.sh      # GUI功能演示
./tests/font_demo.sh     # 字体功能演示
```

**测试文件说明：**
- `tests/test_data/` - 测试轨迹文件（GPX、KML、OVJSN格式）
- `tests/*.md` - 详细测试文档
- `tests/*.sh` - 测试脚本

目录结构示例：

```bash
├── main (或 main.exe)  // 主程序
├── config.ini          // 配置文件
├── build.sh            // Linux/macOS构建脚本
├── build.bat           // Windows构建脚本
└── source_data         // 三方路径数据存放目录，当前支持 KML、GPX、Ovjsn 格式数据
```

### 5.2 运行方式

工具支持两种运行方式：

#### GUI模式（推荐）
双击运行 `main` 或 `main.exe`，启动图形界面进行操作。

#### 命令行模式
在终端中运行 `./main --cli` 使用命令行模式。

### 5.3 GUI界面使用

启动程序后，将打开图形用户界面：

1. **选择文件目录**：
   - 点击"源文件目录"右侧的"选择目录"按钮
   - 选择包含轨迹文件的目录（支持KML、GPX、Ovjsn格式）

2. **设置输出目录**：
   - 点击"输出目录"右侧的"选择目录"按钮
   - 选择CSV文件保存目录（可选，默认保存到源文件目录）

3. **配置参数**：
   - **时间设置**：
     - 开始时间：设置轨迹起始时间（格式：2024-01-01 08:00:00）
     - 结束时间：设置轨迹结束时间（可选，用于精确时间分配）
   - **海拔设置**：设置默认海拔高度（单位：米）
   - **速度设置**：
     - 选择"自动计算"或"手动指定"
     - 手动指定时可设置固定速度值（m/s）
   - **插点设置**：
     - 启用/禁用轨迹插点功能
     - 设置插点距离阈值（米）

4. **开始处理**：
   - 点击"开始处理"按钮执行转换
   - 界面会显示处理进度和状态信息

5. **保存配置**：
   - 点击"保存配置"按钮将当前设置保存到config.ini文件

**批量处理特性：**
- GUI界面支持批量处理，会自动扫描目录内所有支持的文件
- 每个源文件会生成对应的CSV文件，文件名格式：`原文件名_steplife.csv`
- 处理过程中会显示实时进度和状态信息


## 6. 数据导入指南

### 6.1 导入奥维互动地图数据

>奥维互动地图，专业地理规划软件
>
>跨平台地图浏览器，支持 iOS (iPhone、iPad)、Android、Windows、macOS、Linux 等流行平台。
>
>拥有强大的设计功能与地理信息展现技术，可满足各行各业地理信息规划的需求。它不仅是您工作上的好帮手，也是您探索未知世界的更好伴侣。

#### 步骤说明：

`Step1.` 你需要下载奥维互动地图应用，下载地址：https://www.ovital.com/download/

`Step2.` 打开奥维互动地图应用，输入起点终点，选择你的路线后并双击，转换成轨迹并保存。具体操作如下图所示

<img width="1485" alt="map_search" src="https://github.com/user-attachments/assets/940807b5-9649-4cb4-96b2-89a80fe327f3" />

<img width="1485" alt="map_search2" src="https://github.com/user-attachments/assets/3c3516aa-841b-44dc-ba6f-7dbb7ce7839a" />

`Step3.` 点击左下角收藏夹，选中你的轨迹进行导出，导出格式请务必选择`Omap json格式`

<img width="1485" alt="export_tracks" src="https://github.com/user-attachments/assets/a4dcc843-efd0-473e-8930-78301e1009c4" />

`Step4.` 导出后的数据放置 `source_data` 目录下，

- `Mac用户`: 在对应目录中打开终端，输入 `./main` 即可。
```bash
./main
```

- `Windows用户`: 在对应目录中打开双击 `main.exe` 或者在终端输入 `./main.exe` 即可。
```bash
./main.exe
```

<img width="1020" alt="Snipaste_2025-05-25_10-46-02" src="https://github.com/user-attachments/assets/03b1187a-d506-40c6-bd43-70aa538c97b2" />

`Step5.` 程序会生成的`ouput.csv`文件，导入手机中，打开选择选择一生足迹，即可完成数据录入



### 6.2 导入 KML、GPX 数据

>KML（Keyhole Markup Language）是一种基于 XML 的文件格式，用于描述地理空间数据。它最初由 Keyhole 公司开发，后被 Google 收购并广泛应用于 Google Earth 等地理信息系统中。KML 文件可以定义点、线、面、模型等地理要素，并支持三维可视化，非常适合用于毕业设计中展示三维立体图。
>
>
>
>GPX（GPS eXchange Format, GPS 交换格式）是一个 XML 格式，为应用软件设计的通用 GPS 数据格式。它可以用来描述路点、轨迹、路程。这个格式是免费的，可以在不需要付任何许可费用的前提下使用。它的标签保存位置，海拔和时间，可以用来在不同的 GPS 设备和软件之间交换数据。如查看轨迹、在照片的exif数据中嵌入地理数据。



`Step1.` 你可以从各类地图软件中获取KML数据，再次不做举例

`Step2.` 导出后的数据放置 `source_data` 目录下，在对应目录中打开终端，输入 `./main` 即可。

```bash
./main
```

`Step3.` 程序将生成 `output.csv`，可在手机中导入至一生足迹 App。



## 7. 注意事项

- 请确保文件格式正确，命名规范，以便程序正常识别
- 当前版本为命令行工具，后续将考虑支持图形界面
- 如有问题或建议，欢迎提交 [Issue](https://github.com/hygao1024/steplife-universal-importer/issues) 或者 [PR](https://github.com/hygao1024/steplife-universal-importer/pulls)



## 8. 致谢

感谢每一位愿意记录人生、珍视记忆的你。
如果你觉得项目有帮助，请点个 Star ⭐️！仓库地址：https://github.com/hygao1024/steplife-universal-importer
