# DSA Spaced Repetition System (dsa\_srs)

A terminal-based spaced repetition tool designed to help you master Data Structures & Algorithms (DSA) concepts and LeetCode problems. Built in Go with a TUI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss).

---

## 🚀 Features

* **Dual card types**: Manage both theory concepts and coding problems.
* **Spaced Repetition Algorithm**: Implements the SM-2 algorithm to optimize your review intervals.
* **Local JSON Storage**: All cards are stored in a `cards.json` file for easy backup and portability.
* **TUI Interface**: Intuitive terminal UI with forms for adding cards and menus for reviewing due items.
* **Automatic scheduling**: The app calculates next review dates, intervals, and tracks ease factors.
* **Progress tracking**: Displays how many cards are due and your position in the review session.

---

## 💻 Installation

Ensure you have Go installed (version 1.18+).

```bash
# Install the binary:
go install github.com/TheChirag356/dsa_srs@latest

# The binary will be installed to $GOPATH/bin or $HOME/go/bin
```

Alternatively, clone and build from source:

```bash
git clone https://github.com/TheChirag356/dsa_srs.git
cd dsa_srs
go build -o dsa_srs
```

---

## ⚙️ Usage

After installation, run the app:

```bash
dsa_srs
```

### Main Menu

* **Review Due Cards**: Start a review session for all cards whose next review date has passed.
* **Add Concept Card**: Enter a new DSA concept with title, notes, and tags.
* **Add Problem Card**: Enter a LeetCode problem with title, link, and topic.
* **Exit**: Quit the application.

### Reviewing Cards

During a review session:

* All due cards are listed one by one.
* View the card details (e.g., concept notes or problem link).
* Rate your recall:

  * `1` = Again (reschedule soon)
  * `2` = Good (standard interval)
  * `3` = Easy (longer interval)

The app updates each card's repetition count, ease factor, and schedules the next review automatically.

### Data File

All cards are stored in `data/cards.json` in the current working directory. You can backup, version-control, or copy this file between machines to sync your progress.

---

## 📝 Configuration

* **Card file path**: By default, the app reads/writes `cards.json`. You can modify the path via the `constants.CardFilePath` in code.

---

## Contributing

Contributions are welcome! Feel free to fork, open issues, or submit pull requests. Ideas for improvements:

* Filtering and searching cards
* Editing and deleting existing cards
* Import/export decks
* Enhanced statistics and progress charts

Please follow the repository's [CONTRIBUTING.md](CONTRIBUTING.md) guidelines.

---

## 📄 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

Happy coding and learning! 🎓
