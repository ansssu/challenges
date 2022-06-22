go test -bench=. 
go test -bench=.  -run=^# 
go test -bench=.  -benchtime=1000x 
go test -bench=.  -benchtime=10s   -benchmem -run=^# 
go test -bench=.  -benchmem
go test -bench=Prime -count 5 | tee  old.txt
benchstat old.txt
benchstat old.txt new.txt 