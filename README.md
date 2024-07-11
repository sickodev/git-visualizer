# Git Vizualizer

Learned from [flaviocopes](https://flaviocopes.com/go-git-contributions/)

Welcome to **Git Vizualizer**, a project that allows you to visualize your local Git repository contributions using the Go programming language. This tool generates a graphical representation of your commit history, helping you track your contributions over time.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Configuration](#configuration)

## Features

- **Visual Representation**: Generate a graphical representation of your Git commit history.
- **Terminal Agnostic**: See the commit history in your favourite terminal.
- **Simple and Fast**: Built with Go for efficient performance.
- **Local Repository Support**: Works with any local Git repository.

## Installation

To install and set up **Git Visualizer**, follow these steps:

1. **Clone the Repository**:

    ```sh
    git clone https://github.com/sickodev/git-visualizer.git
    cd git-visualizer
    ```

2. **Install Dependencies**:
    Ensure you have Go installed on your machine. If not, download and install it from [golang.org](https://golang.org/dl/).

3. **Build the Project**:

    ```sh
    go build
    ```

4. **Run the Executable**:

    ```sh
    ./git-visualizer
    ```

## Usage

To use the tool, navigate to your local Git repository and run the executable. The tool will process the commit history and generate a visualization.

1. **Navigate to Your Local Git Repository**:

    ```sh
    cd /path/to/your/git/repository
    ```

2. **Run the Tool**:

    ```sh
    /path/to/git-visualizer
    ```

3. **View the Visualization**:
    The tool will generate a graphical representation of your commit history, which can be viewed on your terminal.

## Configuration

You can customize the visualization by modifying the configuration file (`config.json`) included in the project. The configuration options include:

---

Thank you for using **Visualize Your Local Git Contributions with Go**! If you have any questions or feedback, please feel free to open an issue on the [GitHub repository](https://github.com/sickodev/git-vizualizer)
