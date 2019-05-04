# 基于 Cookie 切换不同后端环境

## 使用方法

checkout 代码后，执行 `docker-compose up -d` 启动

访问 `http://127.0.0.1:9000/envs` 查看后端环境列表及切换

访问 `http://127.0.0.1:9000/` 可以查看到当前使用的环境

执行 `docker-compose logs --tail=10 -f` 可以观察请求的路径，访问到了哪个对应的环境
