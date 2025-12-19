#!/bin/bash

# 输出文件验证脚本

echo "=== 一生足迹数据导入器 - 输出验证 ==="
echo

# 获取项目根目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
OUTPUT_DIR="$PROJECT_ROOT/output"

if [ ! -d "$OUTPUT_DIR" ]; then
    echo "❌ 输出目录不存在：$OUTPUT_DIR"
    echo "请先运行测试生成输出文件"
    exit 1
fi

echo "📂 扫描输出目录：$OUTPUT_DIR"
echo

# 查找所有CSV文件
CSV_FILES=$(find "$OUTPUT_DIR" -name "*.csv" -type f)

if [ -z "$CSV_FILES" ]; then
    echo "❌ 未找到CSV文件"
    exit 1
fi

echo "📋 发现的CSV文件："
echo "$CSV_FILES" | while read -r file; do
    echo "  $(basename "$file")"
done
echo

# 验证每个CSV文件
echo "🔍 验证CSV文件内容："
echo

for csv_file in $CSV_FILES; do
    echo "📄 验证文件：$(basename "$csv_file")"

    # 检查文件是否存在且不为空
    if [ ! -s "$csv_file" ]; then
        echo "  ❌ 文件为空"
        continue
    fi

    # 读取文件头
    header=$(head -n 1 "$csv_file")
    expected_header="dataTime,locType,longitude,latitude,heading,accuracy,speed,distance,isBackForeground,stepType,altitude"

    if [ "$header" != "$expected_header" ]; then
        echo "  ❌ 文件头不正确"
        echo "    期望：$expected_header"
        echo "    实际：$header"
        continue
    fi

    echo "  ✅ 文件头正确"

    # 统计数据行数
    data_lines=$(tail -n +2 "$csv_file" | wc -l)
    echo "  📊 数据行数：$data_lines"

    # 检查数据格式
    valid_lines=0
    invalid_lines=0

    tail -n +2 "$csv_file" | while IFS=',' read -r dataTime locType longitude latitude heading accuracy speed distance isBackForeground stepType altitude; do
        # 验证必需字段
        if [[ -n "$dataTime" && -n "$longitude" && -n "$latitude" ]]; then
            valid_lines=$((valid_lines + 1))
        else
            invalid_lines=$((invalid_lines + 1))
        fi
    done

    echo "  ✅ 有效数据行：$valid_lines"
    if [ $invalid_lines -gt 0 ]; then
        echo "  ⚠️  无效数据行：$invalid_lines"
    fi

    # 显示前几行数据作为示例
    echo "  📝 数据示例："
    head -n 3 "$csv_file" | while read -r line; do
        echo "    $line"
    done

    echo
done

echo "=== 验证完成 ==="
echo
echo "💡 提示："
echo "  - 检查时间戳是否在配置的时间范围内"
echo "  - 检查经纬度坐标是否正确"
echo "  - 检查海拔高度是否为配置的默认值"
echo "  - 检查速度值是否合理"

