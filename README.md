# QUERY_SEARCH

A lightweight command-line application that leverages [USearch](https://github.com/unum-cloud/usearch) for efficient vector similarity search based on user queries.  
It also uses the **OpenAI Embedding API** to generate vector representations of input texts.

## Features

- Embedding input text data
- Indexing vectors for fast retrieval
- Performing nearest-neighbor similarity searches

## Installation

You can either **build `query_search` from source** manually or **download a prebuilt artifact** from GitHub Actions.

## Building from Source

1. **Clone the repository:**

    ```bash
    git clone https://github.com/weisifan/query_search.git
    cd query_search
    ```

2. **Install Go:**

    ```bash
    sudo apt update
    sudo apt install golang-go
    ```

    Ensure you have **Go 1.18+**:

    ```bash
    go version
    ```

3. **Install or extract USearch C library:**

    You need the shared library `libusearch_c.so`.

    **Option A: Download the prebuilt USearch .deb package:**

    ```bash
    wget https://github.com/unum-cloud/usearch/releases/download/v2.17.7/usearch_linux_amd64_2.17.7.deb

    ar x usearch_linux_amd64_2.17.7.deb
    
    tar --use-compress-program=unzstd -xf data.tar.zst
    ```

    You will find:
    - `usr/local/lib/libusearch_c.so`
    - `usr/local/include/usearch/*.hpp`

    **Set your environment:**

    ```bash
    export LD_LIBRARY_PATH=$(pwd)/usr/local/lib:$LD_LIBRARY_PATH
    ```

    **Option B: Build USearch from source:**  
    Clone the [USearch GitHub repository](https://github.com/unum-cloud/usearch) and build using CMake.

4. **Build `query_search`:**

    ```bash
    go build -o query_search
    ```

    After building, you will have a local executable `query_search`.

## Using Prebuilt Artifact

If you prefer not to build manually:

1. Go to this repository → **Actions** tab.
2. Select the latest successful workflow.
3. Download the artifact (.zip file).
4. Unzip the file — you will get:
    - `query_search`
    - `libusearch_c.so`
    - `run.sh`

**Example to run:**

```bash
./run.sh <api-key> <path-to-text-file>
```

## Usage

**Basic command structure:**

```bash
./query_search <api-key> <path-to-text-file>
```

- `<api-key>`: Your API key for the embedding service (e.g., OpenAI API key).
- `<path-to-text-file>`: A plain text file where each line is a separate text entry to be indexed.

**Example:**

```bash
./query_search <your-openai-api-key> demo.txt
```

This will:
- Embed each line from `demo.txt`
- Add the embeddings to the USearch index
- Enable similarity search queries through a REPL interface

## Environment Setup Reminder

Before running `query_search`, ensure the shared library path is set correctly:

```bash
export LD_LIBRARY_PATH=/path/to/libusearch_c_folder:$LD_LIBRARY_PATH
```

Otherwise, you may encounter errors like:

```
error while loading shared libraries: libusearch_c.so: cannot open shared object file
```

## License

Currently, there is no specific license attached to this project.

## Notes

- `libusearch_c.so` must be accessible at runtime (either installed globally or via `LD_LIBRARY_PATH`).
- A helper script `run.sh` can automate environment setup and running.
- Future improvements may include static linking to eliminate `.so` dependencies for full portability.

## Quick Commands Recap

| Task | Command |
|:-----|:--------|
| Clone project | `git clone https://github.com/weisifan/query_search.git` |
| Build binary | `go build -o query_search ./cmd` |
| Set library path | `export LD_LIBRARY_PATH=$(pwd)/usr/local/lib:$LD_LIBRARY_PATH` |
| Run | `./query_search <api-key> <text-file>` |
