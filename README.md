QUERY SEARCH

A command-line application that leverages USearch (https://github.com/unum-cloud/usearch) for efficient vector similarity search based on user queries. It also leverages the OpenAI Embedding API to generate vector representations of input texts.

This project demonstrates:
- Embedding input data
- Indexing vectors
- Performing fast nearest-neighbor searches

------------------------------------------------------------

INSTALLATION

You can either build query_search from source manually or download a prebuilt artifact from GitHub Actions.

------------------------------------------------------------

BUILDING FROM SOURCE

1. Clone the repository:

git clone https://github.com/weisifan/query_search.git
cd query_search


2. Install Go:

sudo apt update
sudo apt install golang-go

Check the Go version (needs Go 1.18+):

go version

3. Install or extract USearch C library:

You need the shared library libusearch_c.so.

Option A: Download prebuilt USearch .deb package and extract manually:

wget https://github.com/unum-cloud/usearch/releases/download/v2.17.7/usearch_linux_amd64_2.17.7.deb
ar x usearch_linux_amd64_2.17.7.deb
tar --use-compress-program=unzstd -xf data.tar.zst

You will find:
usr/local/lib/libusearch_c.so
usr/local/include/usearch/*.hpp

Set your environment:

export LD_LIBRARY_PATH=$(pwd)/usr/local/lib:$LD_LIBRARY_PATH

Option B: Build USearch from source:

Clone the USearch GitHub repository and build it using CMake.

4. Build the query_search binary:

go build -o query_search

Now you have a local executable "query_search".

------------------------------------------------------------

USING PREBUILT ARTIFACT

If you don't want to build manually:

1. Go to this repository -> Actions tab.
2. Select the latest successful workflow.
3. Download the artifact (ZIP file).
4. Unzip it, you will get:
   - query_search (the executable)

You still need libusearch_c.so separately, either from building manually or from the extracted .deb package.

Make sure to set:

export LD_LIBRARY_PATH=/path/to/your/usr/local/lib:$LD_LIBRARY_PATH

Example:

export LD_LIBRARY_PATH=$(pwd)/usr/local/lib:$LD_LIBRARY_PATH

------------------------------------------------------------

USAGE

Basic command structure:

./query_search <api-key> <path-to-text-file>

- <api-key>: Your API key for the embedding service (e.g., OpenAI, or any compatible service).
- <path-to-text-file>: A plain text file where each line is a separate text to be indexed.

Example:

./query_search <your-openai-api-key> demo.txt

This will:
- Embed each line from demo.txt
- Add it to the USearch index
- Allow you to perform nearest neighbor search queries.

------------------------------------------------------------

ENVIRONMENT SETUP SUMMARY

Before running, always make sure:

export LD_LIBRARY_PATH=/path/to/libusearch_c_folder:$LD_LIBRARY_PATH

Otherwise you might see errors like:

error while loading shared libraries: libusearch_c.so: cannot open shared object file

------------------------------------------------------------

LICENSE

None

------------------------------------------------------------

NOTES

- libusearch_c.so must be accessible at runtime (either installed globally, or via LD_LIBRARY_PATH).
- If needed, you can create a helper script run.sh to automate environment setup.
- Future improvements may include static linking for full portability without .so dependencies.

------------------------------------------------------------

QUICK COMMANDS RECAP

Clone project:
git clone https://github.com/weisifan/query_search.git

Build binary:
go build -o query_search

Set library path:
export LD_LIBRARY_PATH=$(pwd)/usr/local/lib:$LD_LIBRARY_PATH

Run:
./query_search <api-key> <text-file>
