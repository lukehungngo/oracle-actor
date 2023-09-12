#!/usr/bin/env bash
# Generate the go binary and places them in their respective directories.

set -eux -o pipefail

command_installed() {
	type "$1" &>/dev/null
}

required_commands=("solc" "abigen")
for cmd in "${required_commands[@]}"; do
	if ! command_installed "$cmd"; then
		echo "$cmd is not installed. Please install it before running this script."
		exit 1
	fi
done

readonly sc_dir="internal/_contracts"
artifacts_dir="${sc_dir}/artifacts"
readonly pkg_dir="pkg/generated/contracts"

if [ ! -d "$artifacts_dir" ]; then
	echo "Creating artifacts directory..."
	mkdir -p "$artifacts_dir"
fi

function camelToSnake() {
	local input="$1"
	local output=""

	if [[ -z "$input" ]]; then
		echo "Error: No input provided."
		return 1
	fi

	output=$(perl -pe 's/(?<=[a-z])(?=[A-Z])/_/g' <<<"$input" | perl -pe 'y/A-Z/a-z/')
	echo "$output"
}

solidity_files=$(find "${sc_dir}/src" -maxdepth 1 -type f -name "*.sol" -exec basename {} .sol \;)

for file in $solidity_files; do
	go_file_name=$(camelToSnake "$(basename "$file" .sol)")
	# gobin_dir_name="$(echo "${file:0:1}" | tr '[:upper:]' '[:lower:]')${file:1}"
	mkdir -p "$pkg_dir/$go_file_name"

	solc --optimize --optimize-runs=200 --base-path "${sc_dir}" \
		--metadata --metadata-literal --abi "${sc_dir}/src/${file}.sol" \
		-o "${artifacts_dir}/abi/" --overwrite

	solc --optimize --optimize-runs=200 --base-path "${sc_dir}" --include-path . \
		--bin "${sc_dir}/src/${file}.sol" \
		-o "${artifacts_dir}/bin/" --overwrite

	abigen --abi "${artifacts_dir}/abi/${file}.abi" \
		--bin "${artifacts_dir}/bin/${file}.bin" \
		--pkg "${go_file_name}" --type "${file}" \
		--out "${pkg_dir}/$go_file_name/${go_file_name}.go"

done
