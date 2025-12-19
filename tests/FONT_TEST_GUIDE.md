# 🎨 自定义字体使用指南

## 📋 字体功能概述

GUI版本的一生足迹数据导入器现在支持自定义字体显示，使用小米字体(MiSans)提供更好的视觉体验。

## 🔤 字体信息

**字体文件**：`./resource/MiSans-Regular.otf`
- **字体名称**：MiSans Regular
- **文件大小**：6.5MB (6,499,984 bytes)
- **类型**：OpenType字体(.otf)
- **设计**：小米公司设计，现代简洁风格

## ⚙️ 字体功能特性

### 自动加载
- 程序启动时自动检测 `./resource/MiSans-Regular.otf` 文件
- 如果字体文件存在且有效，则自动加载并应用
- 如果字体文件不存在或损坏，则使用系统默认字体

### 应用范围
- ✅ 所有UI文本标签
- ✅ 按钮文字
- ✅ 输入框文字
- ✅ 菜单和对话框
- ✅ 状态信息显示

### 兼容性
- ✅ **macOS**：原生支持OpenType字体
- ✅ **Windows**：支持OpenType字体
- ✅ **Linux**：支持OpenType字体（需要字体渲染支持）

## 🧪 字体测试

### 验证字体加载

运行程序时查看启动日志：
```bash
./main
```

**成功日志示例**：
```
{"level":"INFO","time":"2025-12-19 18:35:13.110","caller":"gui/main.go:63","message":"成功加载自定义字体: MiSans-Regular.otf (大小: 6499984 bytes)"}
```

**失败日志示例**：
```
{"level":"INFO","time":"2025-12-19 18:35:13.110","caller":"gui/main.go:63","message":"字体文件不存在 ./resource/MiSans-Regular.otf，使用系统默认字体"}
```

### 视觉验证

启动GUI界面后，观察：
- [ ] 字体是否更清晰圆润
- [ ] 中英文混排是否协调
- [ ] 界面整体美观度是否提升
- [ ] 与系统默认字体对比差异

## 🔧 字体管理

### 添加新字体

1. 将字体文件放置在 `./resource/` 目录下
2. 修改 `internal/gui/main.go` 中的字体文件名
3. 重新构建程序：`go build -o main ./cmd`

### 支持的字体格式

- ✅ **.otf** (OpenType Font) - 推荐
- ✅ **.ttf** (TrueType Font)
- ⚠️ **.woff/.woff2** - Web字体，需要额外处理

### 字体大小调整

如果需要调整字体大小，可以修改主题：

```go
func (c *customTheme) Size(name fyne.ThemeSizeName) float32 {
    switch name {
    case theme.SizeNameText:
        return 14 // 自定义文字大小
    case theme.SizeNameHeadingText:
        return 18 // 自定义标题大小
    default:
        return theme.DefaultTheme().Size(name)
    }
}
```

## 🐛 字体问题排查

### 问题1：字体显示为系统默认字体

**可能原因**：
- 字体文件不存在或损坏
- 字体格式不支持
- 系统缺少字体渲染支持

**解决方案**：
```bash
# 检查字体文件
ls -la ./resource/MiSans-Regular.otf

# 验证文件完整性
file ./resource/MiSans-Regular.otf

# 重新下载字体文件
curl -o ./resource/MiSans-Regular.otf [字体下载地址]
```

### 问题2：字体显示异常/乱码

**可能原因**：
- 字体文件编码问题
- 系统字体渲染器不支持

**解决方案**：
- 尝试其他字体格式(.ttf)
- 检查系统字体设置
- 更新图形驱动

### 问题3：性能影响

**字体大小**：6.5MB可能对启动时间有轻微影响
**内存占用**：加载后会占用额外内存
**渲染性能**：OpenType字体渲染通常比系统默认字体更流畅

## 📊 字体对比

| 特性 | 系统默认字体 | MiSans字体 |
|-----|-------------|-----------|
| **文件大小** | 0KB (系统内置) | 6.5MB |
| **加载时间** | 即时 | < 0.1秒 |
| **视觉效果** | 一般 | 优秀 |
| **中英文支持** | 良好 | 优秀 |
| **可读性** | 良好 | 优秀 |

## 🎨 字体设计特色

MiSans字体具有以下特色：
- **现代简洁**：线条流畅，结构清晰
- **中英文优化**：针对中文和英文的专门优化
- **多场景适用**：适合界面显示和阅读
- **品牌一致性**：小米生态统一字体

## 📝 开发说明

### 代码位置
- **字体加载**：`internal/gui/main.go:48-63`
- **主题实现**：`internal/gui/main.go:65-85`
- **应用设置**：`internal/gui/main.go:99-101`

### 扩展功能
可以轻松添加更多字体变体：
```go
// 添加粗体字体
g.fontBold = fyne.NewStaticResource("MiSans-Bold", boldFontData)

// 在主题中使用
func (c *customTheme) Font(style fyne.TextStyle) fyne.Resource {
    if style.Bold {
        return c.fontBold
    }
    return c.fontRegular
}
```

---

**🎯 总结**：自定义字体功能让GUI界面更加美观专业，提升了用户体验！
