#!/bin/bash

# 定义镜像版本号，方便修改
VERSION="0.0.1"
IMAGE_NAME="number-sender"
FULL_IMAGE_NAME="${IMAGE_NAME}:${VERSION}"
TAR_FILE="number-sender.${VERSION}.tar"

# 显示帮助信息
show_help() {
    echo "使用方法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  b    执行构建Docker镜像功能"
    echo "  d    执行停止Docker容器功能"
    echo "  l    执行查看容器日志功能"
    echo "  m    执行镜像迁移功能"
    echo "  u    执行启动Docker容器功能"
    echo "  h    显示此帮助信息"
    echo ""
    echo "如果不提供任何选项，将进入交互式选择模式。"
}

# 构建Docker镜像 (docker-build.sh)
build_image() {
    # 显示脚本开始执行
    echo "========================================"
    echo "Docker镜像编译脚本开始执行"
    echo "镜像名称：${IMAGE_NAME}"
    echo "镜像版本：${VERSION}"
    echo "完整镜像名：${FULL_IMAGE_NAME}"
    echo "========================================"

    # 检查镜像tag是否已存在
    echo "\n1. 检查镜像 ${FULL_IMAGE_NAME} 是否已存在..."
    if docker image inspect "${FULL_IMAGE_NAME}" > /dev/null 2>&1; then
        echo "   ✓ 镜像 ${FULL_IMAGE_NAME} 已存在"
        
        # 检查该tag是否正在被运行
        echo "2. 检查镜像 ${FULL_IMAGE_NAME} 是否正在被运行..."
        # 使用awk精确匹配Image字段（第二列）检查容器是否正在运行
        RUNNING_CONTAINERS=$(docker ps --format "{{.ID}} {{.Image}} {{.Names}}" | awk -v img="${FULL_IMAGE_NAME}" '$2 == img')
        if [ -n "${RUNNING_CONTAINERS}" ]; then
            echo "   ✗ 错误：镜像 ${FULL_IMAGE_NAME} 正在被以下容器运行："
            echo "${RUNNING_CONTAINERS}" | awk '{print "   " $1 " " $3}'
            echo "   无法编译，请先停止相关容器"
            return 1
        else
            echo "   ✓ 镜像 ${FULL_IMAGE_NAME} 未在运行"
            
            # 询问是否删除该空置tag并继续编译
            echo "3. 镜像 ${FULL_IMAGE_NAME} 存在但未运行"
            read -p "   是否删除该空置镜像并继续编译？(y/n，默认n) " -n 1 -r
            echo
            
            if [[ $REPLY =~ ^[Yy]$ ]]; then
                echo "4. 删除镜像 ${FULL_IMAGE_NAME}..."
                if docker rmi "${FULL_IMAGE_NAME}"; then
                    echo "   ✓ 镜像 ${FULL_IMAGE_NAME} 删除成功"
                else
                    echo "   ✗ 删除镜像失败，退出编译"
                    return 1
                fi
            else
                echo "   ✗ 取消编译"
                return 0
            fi
        fi
    else
        echo "   ✓ 镜像 ${FULL_IMAGE_NAME} 不存在，可以开始编译"
    fi

    # 构建Docker镜像
    echo "\n5. 开始构建Docker镜像：${FULL_IMAGE_NAME}"
    if docker build -t "${FULL_IMAGE_NAME}" .; then
        echo "   ✓ Docker镜像构建成功：${FULL_IMAGE_NAME}"
        
        # 更新docker-compose.yml中的镜像版本号
        echo "6. 更新docker-compose.yml中的镜像版本号为：${VERSION}"
        if sed -i "s|image: ${IMAGE_NAME}:.*|image: ${FULL_IMAGE_NAME}|g" docker-compose.yml; then
            echo "   ✓ docker-compose.yml更新成功"
        else
            echo "   ✗ docker-compose.yml更新失败"
            return 1
        fi
        
        echo "\n========================================"
        echo "Docker镜像编译脚本执行完成"
        echo "镜像名称：${FULL_IMAGE_NAME}"
        echo "编译状态：成功"
        echo "========================================"
    else
        echo "   ✗ Docker镜像构建失败"
        echo "\n========================================"
        echo "Docker镜像编译脚本执行完成"
        echo "编译状态：失败"
        echo "========================================"
        return 1
    fi
}

# 启动Docker容器 (docker-up.sh)
start_containers() {
    echo "启动Docker容器服务..."
    
    # 覆盖拷贝配置文件
    echo "1. 覆盖拷贝配置文件..."
    SOURCE_CONFIG="./config/config.toml"
    TARGET_DIR="./docker/app/config"
    TARGET_CONFIG="${TARGET_DIR}/config.toml"
    
    # 检查源文件是否存在
    if [ -f "${SOURCE_CONFIG}" ]; then
        # 确保目标目录存在
        mkdir -p "${TARGET_DIR}"
        # 覆盖拷贝配置文件
        if cp -f "${SOURCE_CONFIG}" "${TARGET_CONFIG}"; then
            echo "   ✓ 配置文件已成功拷贝到 ${TARGET_CONFIG}"
        else
            echo "   ✗ 配置文件拷贝失败"
            return 1
        fi
    else
        echo "   ✗ 错误：源配置文件 ${SOURCE_CONFIG} 不存在"
        return 1
    fi
    
    # 设置docker目录权限
    echo "2. 设置docker目录权限..."
    if chmod 777 -R ./docker; then
        echo "   ✓ docker目录权限设置成功"
    else
        echo "   ✗ docker目录权限设置失败"
        return 1
    fi
    
    docker compose up -d
    echo "Docker容器服务启动完成"

    echo "查看容器状态："
    docker compose ps
}

# 停止Docker容器 (docker-down.sh)
stop_containers() {
    echo "停止Docker容器服务..."
    docker compose down
    echo "Docker容器服务停止完成"

    echo "查看当前运行的容器："
    docker ps
}

# 查看容器日志 (docker-log.sh)
show_logs() {
    echo "查看number-sender容器日志..."
    docker logs number-sender
}

# 导出镜像
export_image() {
    echo "========================================"
    echo "开始导出镜像: ${FULL_IMAGE_NAME}"
    echo "导出文件: ${TAR_FILE}"
    echo "========================================"
    
    if docker save -o "${TAR_FILE}" "${FULL_IMAGE_NAME}"; then
        echo "✓ 镜像导出成功: ${TAR_FILE}"
        echo "导出文件大小: $(du -h "${TAR_FILE}" | cut -f1)"
    else
        echo "✗ 镜像导出失败"
        return 1
    fi
}

# 导入镜像
import_image() {
    echo "========================================"
    echo "开始导入镜像: ${TAR_FILE}"
    echo "========================================"
    
    if [ ! -f "${TAR_FILE}" ]; then
        echo "✗ 错误: 导出文件 ${TAR_FILE} 不存在"
        return 1
    fi
    
    if docker load -i "${TAR_FILE}"; then
        echo "✓ 镜像导入成功"
        echo "导入的镜像: $(docker images | grep "${IMAGE_NAME}" | grep "${VERSION}")"
    else
        echo "✗ 镜像导入失败"
        return 1
    fi
}

# 镜像迁移交互逻辑
migration_menu() {
    echo "========================================"
    echo "Docker镜像迁移功能"
    echo "当前配置:"
    echo "  镜像名称: ${IMAGE_NAME}"
    echo "  镜像版本: ${VERSION}"
    echo "  完整镜像名: ${FULL_IMAGE_NAME}"
    echo "  导出文件名: ${TAR_FILE}"
    echo "========================================"
    echo "请选择要执行的操作:"
    echo "1) 导出镜像 (docker save)"
    echo "2) 导入镜像 (docker load)"
    echo "直接回车退出镜像迁移功能"
    echo ""
    
    read -p "请输入您的选择 [1/2]: " -r
    echo
    
    if [ -z "$REPLY" ]; then
        # 默认回车退出
        echo "退出镜像迁移功能"
        return 0
    fi
    
    case "$REPLY" in
        1|export)
            export_image
            ;;
        2|import)
            import_image
            ;;
        *)
            echo "无效选择: $REPLY"
            return 1
            ;;
    esac
}

# 执行相应的功能
run_script() {
    case "$1" in
        b)
            echo "执行构建Docker镜像功能..."
            build_image
            ;;
        d)
            echo "执行停止Docker容器功能..."
            stop_containers
            ;;
        l)
            echo "执行查看容器日志功能..."
            show_logs
            ;;
        m)
            echo "执行镜像迁移功能..."
            migration_menu
            ;;
        u)
            echo "执行启动Docker容器功能..."
            start_containers
            ;;
        *)
            echo "无效选项: $1"
            show_help
            exit 1
            ;;
    esac
}

# 主程序
if [ $# -eq 0 ]; then
    # 没有参数，进入交互式选择
    echo "========================================"
    echo "Docker 工具脚本"
    echo "========================================"
    echo "请选择要执行的操作:"
    echo "b) 构建Docker镜像 (docker build)"
    echo "d) 停止Docker容器 (docker compose down)"
    echo "l) 查看容器日志 (docker logs)"
    echo "m) 镜像迁移 (docker save/load)"
    echo "u) 启动Docker容器 (docker compose up -d)"
    echo "直接回车退出脚本"
    echo ""
    
    read -p "请输入您的选择 [b/d/l/m/u]: " -n 1 -r
    echo
    
    if [ -z "$REPLY" ]; then
        # 默认回车退出
        echo "退出脚本"
        exit 0
    elif [[ $REPLY =~ ^[bBdDlLmMuU]$ ]]; then
        # 执行选择的功能
        run_script "${REPLY,,}"
    else
        echo "无效选择: $REPLY"
        exit 1
    fi
elif [ $# -eq 1 ]; then
    # 有一个参数，直接执行
    case "$1" in
        h|--help)
            show_help
            exit 0
            ;;
        *)
            run_script "$1"
            ;;
    esac
else
    # 参数过多
    echo "错误: 参数过多"
    show_help
    exit 1
fi
