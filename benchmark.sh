# BENCHMARK 1: General, realistic, single file
COMMAND="will test_files/sherlock-holm.txt"
hyperfine --warmup 3 -N -u millisecond "./go-grep $COMMAND" "grep $COMMAND" "rg $COMMAND" --export-markdown g-benchmark.md

# BENCHMARK 2: Large single file
# Generate 1GB~ random file
openssl rand -out sample.txt -base64 $(( 2**30 * 3/4 ))

COMMAND="a4 sample.txt"
hyperfine --warmup 3 -N -u millisecond "./go-grep $COMMAND" "grep $COMMAND" "rg $COMMAND" --export-markdown lf-benchmark.md

rm ./sample.txt
