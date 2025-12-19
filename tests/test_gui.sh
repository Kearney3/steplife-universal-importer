#!/bin/bash

# GUI界面功能验证脚本

echo "🎨 一生足迹数据导入器 - GUI界面测试"
echo "======================================"
echo

# 检查可执行文件
if [ ! -f "./main" ]; then
    echo "❌ 错误：找不到可执行文件 main"
    echo "请先运行以下命令构建程序："
    echo "  go build -o main ./cmd"
    exit 1
fi

echo "✅ 找到可执行文件：main"
echo

# 检查测试数据
if [ ! -d "./test_data" ]; then
    echo "❌ 错误：找不到测试数据目录 test_data"
    exit 1
fi

echo "✅ 找到测试数据目录：test_data"
echo

# 检查配置文件
if [ ! -f "./config.ini" ]; then
    echo "❌ 错误：找不到配置文件 config.ini"
    exit 1
fi

echo "✅ 找到配置文件：config.ini"
echo

echo "🚀 GUI测试指南："
echo "=================="
echo
echo "📋 在图形界面环境中执行以下步骤："
echo
echo "1️⃣ 启动GUI界面："
echo "   ./main"
echo "   或双击 main 可执行文件"
echo
echo "2️⃣ 界面检查："
echo "   □ 窗口标题：'一生足迹数据导入器'"
echo "   □ 界面布局：卡片式设计，800x600像素"
echo "   □ 文件选择区域：源目录和输出目录选择"
echo "   □ 参数设置卡片：时间、海拔、速度、插点设置"
echo "   □ 状态显示卡片：进度条和状态信息"
echo "   □ 操作按钮：开始处理、保存配置"
echo
echo "3️⃣ 功能测试："
echo "   □ 点击'源文件目录'的'选择目录'按钮"
echo "   □ 选择 test_data/ 文件夹"
echo "   □ 设置参数："
echo "     • 开始时间：2024-01-01 08:00:00"
echo "     • 结束时间：2024-01-01 09:00:00"
echo "     • 海拔高度：200.5米"
echo "     • 速度模式：自动计算"
echo "     • 插点距离：50米"
echo "   □ 点击'保存配置'按钮"
echo "   □ 点击'开始处理'按钮"
echo "   □ 观察进度条和状态变化"
echo
echo "4️⃣ 结果验证："
echo "   □ 处理完成后检查 output/ 目录"
echo "   □ 验证生成 sample_track_steplife.csv 等文件"
echo "   □ 检查CSV文件格式和数据内容"
echo
echo "📚 详细测试文档：./GUI_TEST_GUIDE.md"
echo
echo "⚠️  注意事项："
echo "   • 需要在有图形界面的环境中运行"
echo "   • macOS原生支持GUI"
echo "   • Linux需要GTK库：sudo apt-get install libgtk-3-dev"
echo "   • Windows需要OpenGL支持"
echo
echo "🎯 成功标准："
echo "   • 界面正常显示，所有元素可操作"
echo "   • 文件处理成功，生成正确的CSV输出"
echo "   • 配置能正确保存和恢复"
echo
echo "✨ GUI测试准备完成！请在图形环境中继续测试。"
