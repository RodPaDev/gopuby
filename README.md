# GoPuby - Terminal EPUB Reader

### Description
GoPuby is a terminal-based EPUB reader designed to allow users to read and interact with EPUB files directly in their terminal. The application supports text and image rendering and provides a seamless, scroll-based reading experience. It is built with Go and designed to be lightweight and highly customizable.

###  TechStack
- **Go**
- **BBolt** (Lightweight k,v DB)
- **YAML**: (Configuration Files)

### Key Features
1. **Opening EPUB Files**: Load EPUB files by specifying the file as a command-line argument or through an interactive command within the application.
2. **Rendering Content**: Display text and images in the terminal, utilizing iTerm2's capabilities on macOS for image rendering. (Other terminals will contain a link to the image)
3. **Scroll-based Navigation**: Navigate through content using simple keyboard commands.
4. **Table of Contents Interaction**: Access and interact with a sidebar displaying the book's table of contents, with functionality to mark sections as read.
5. **Command Bar**: Use a command bar for additional functionality such as searching text and jumping to sections.
6. **Customization**: Adjust visual settings like font size, color, and family through a configuration file.

### Configuration
- Configuration settings such as `FontSize`, `FontColor`, and `FontFamily` will be managed through a YAML file, allowing users to customize their reading experience based on personal preferences.

### State Management
- **BookSchema**:
  - **id**: Hash of the book as the unique identifier.
  - **filePath**: Path to the EPUB file.
  - **completedSections**: Array of sections marked as read.
  - **currentPos**: Last read position in the book to resume reading.
  - **createdAt**: Timestamp when the user first opened the book.
  - **updatedAt**: Most recent timestamp when the user interacted with the book.
  - **finishedAt**: Timestamp when the user finished reading the book.
  - **config**: Custom settings such as font size or color specific to each book.

### User Data Storage Locations
- **Windows**: Stored in `C:\Users\<username>\AppData\Local\GoPuby\`
- **macOS/Linux**:  Stored in `~/.local/share/GoPuby/` or `~/.config/GoPuby/`

## Usage

To open an epub run
```bash
gopuby <filepath>
```
Alternatively you can run `gopuby` and once open hit `space` and type `:open <filepath>`

### Hotkeys

#### Opening and Loading EPUB Files
| Hotkey | Action                  | Description                                          |
|--------|-------------------------|------------------------------------------------------|
| Space  | Open Command Bar        | Activate the command bar to input commands.          |
| `:open <filepath>` | Open File | Open an EPUB file specified by the file path.    |
| `:list` | List Books in Library | Lists books in library that have been previously opened |

#### Reading and Navigation
| Hotkey | Action                  | Description                                          |
|--------|-------------------------|------------------------------------------------------|
| ↑              | Scroll Up               | Scroll up through the EPUB content.          |
| ↓              | Scroll Down             | Scroll down through the EPUB content.        |
| →              | Navigate Forward        | Go to the next section.                      |
| ←              | Navigate Backward       | Go to the previous section.                  |
| SHIFT + →      | Navigate Forward        | Go to the next section.                      |
| SHIFT + ←      | Navigate Backward       | Go to the previous section.                  |

#### Table of Contents Interaction
| Hotkey | Action                  | Description                                          |
|--------|-------------------------|------------------------------------------------------|
| o      | Toggle Sidebar          | Show or hide the sidebar with the table of contents. |
| p      | Focus on Sidebar        | Switch focus to the sidebar.                         |
| ↑/↓    | Navigate Sidebar        | Move up or down in the table of contents.            |
| c      | Mark as Read/Unread     | Toggle the read status of the selected section.      |

#### Command Bar Features
| Hotkey | Action                  | Description                                          |
|--------|-------------------------|------------------------------------------------------|
| Space  | Open Command Bar        | Activate the command bar to input commands.          |
| `sx`   | Jump to Section         | Input `sx` followed by the section number or name.   |
| `/`    | Search in Page          | Search for text within the current page.             |
| `//`   | Search in Whole Book    | Search for text across the entire book.              |


