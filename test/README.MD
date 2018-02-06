### 测试覆盖率
1. go test . 基本测试
2. go test -coverprofile=c.out 输出测试覆盖文件
3. go tool cover -html=c.out 使用网页查看测试覆盖文件

### 性能测试
1. go test -bench . 基本性能测试
2. go test -bench . -cpuprofile=cpu.out 输出性能测试文件
3. go tool pprof cpu.out 分析性能测试文件，进入交互式命令行
4. 进入性能测试文件交互式命令行后，输入web就可以来到web界面
