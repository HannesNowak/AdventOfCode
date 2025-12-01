part1:
	go build -o part1 part1.go

part2:
	go build -o part2 part2.go

.PHONY: run1 run2 clean

run1: part1
	./part1 <input

run2: part2
	./part2 <input

clean:
	rm -f part1 part2
