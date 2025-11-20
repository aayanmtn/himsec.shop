# himsec.shop

`himsec.shop` is a terminal-based e-commerce application that runs over SSH. It provides a text-based user interface (TUI) for browsing products, adding them to a wishlist, and proceeding to checkout.

## Features

-   Browse a list of products.
-   View detailed information for each product.
-   Add products to a wishlist.
-   A checkout process to purchase items.
-   Accessible from anywhere via SSH.

## Built With

-   [Go](https://golang.org/)
-   [Bubble Tea](https://github.com/charmbracelet/bubbletea)
-   [Lipgloss](https://github.com/charmbracelet/lipgloss)
-   [Wish](https://github.com/charmbracelet/wish)

## Getting Started

### Prerequisites

-   Go 1.21 or later.
-   An SSH client.

### Installation

1.  Clone the repository:
    ```sh
    git clone <repository-url>
    cd himsec.shop
    ```
2.  Generate an SSH host key:
    ```sh
    ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""
    ```
3.  Build the application:
    ```sh
    go build
    ```

## Usage

1.  Run the server:
    ```sh
    go run main.go
    ```
    The server will start on `0.0.0.0:2222` by default.

2.  Connect to the application using an SSH client:
    ```sh
    ssh localhost -p 2222
    ```

Once connected, you can navigate the application using the keyboard:
-   **Up/Down arrows or k/j**: Navigate product lists.
-   **Enter**: View product details.
-   **w**: Add a product to your wishlist.
-   **b**: Go back or proceed to checkout from the main view.
-   **p**: Proceed to payment from checkout.
-   **q** or **Ctrl+c**: Quit the application.
