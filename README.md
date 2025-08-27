<p align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&height=300&color=0:FF1493,100:FFFFFF&text=Brandlovers&fontColor=FFFFFF&animation=fadeIn">
  <h1 align="center">Code Academy - Learning Repository</h1>
</p>

<p align="center">
  <a href="#-what-is-this-repository">About</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-repository-structure">Structure</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-getting-started-for-new-programmers">Getting Started</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-submitting-your-work">Submit Work</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-troubleshooting-common-issues">Help</a>
</p> 

<p align="center">
  <a href="https://github.com/brandlovers-team/code-academy"><img alt="GitHub" src="https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white"></a>
  <a href="https://brandlovrs.slack.com/archives/C09BM1GEUUR"><img alt="Slack" src="https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white"></a>
  <a href="https://go.dev"><img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=blue"></a>
  <a href="https://github.com/cucumber/godog"><img alt="Godog" src="https://img.shields.io/badge/Godog-32B643?style=for-the-badge&logo=cucumber&logoColor=white"></a>
  <a href="https://www.postgresql.org/"><img alt="PostgresSQL" src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white"></a>
  <a href="https://gorm.io/index.html/"><img alt="GORM" src="https://img.shields.io/badge/GORM-316192?style=for-the-badge&logo=go&logoColor=white"></a>
  <a href="https://redis.io/"><img alt="Redis" src="https://img.shields.io/badge/redis-DC382D?style=for-the-badge&logo=redis&logoColor=FFFFFF"></a>
</p>

<p align = "center">
<b> 🎓 Code Academy  | Brandlovers 📚 </b>
</p>


Welcome to the Code Academy learning repository! This is your space to learn, practice, and grow as a programmer. 

## 📚 What is this repository?

This repository is designed as a collaborative learning environment where:
- **Instructors** provide lessons and exercises
- **Students** submit their work for review and feedback
- **Everyone** learns together through code!

## 🗂️ Repository Structure

The repository follows this organized structure:

```
code-academy/
├── README.md           # This file - your guide to the repository
├── .gitignore         # Files that Git should ignore
└── [student-name]/    # Your personal folder (replace with your actual name)
    ├── lesson-01/     # Folder for lesson 1 exercises
    ├── lesson-02/     # Folder for lesson 2 exercises
    └── ...           # More lessons as you progress
```

### Example Structure:
```
code-academy/
├── john-doe/
│   ├── lesson-01/
│   │   ├── hello-world.js
│   │   └── variables.js
│   └── lesson-02/
│       ├── functions.js
│       └── arrays.js
└── jane-smith/
    ├── lesson-01/
    │   └── my-first-program.py
    └── lesson-02/
        └── learning-loops.py
```

## 🚀 Getting Started for New Programmers

### Step 1: Install Git
Git is a version control system that tracks changes in your code. Think of it as a "save game" system for your code!

- **Windows**: Download from [git-scm.com](https://git-scm.com/download/win)
- **Mac**: Install via Terminal: `brew install git` (requires [Homebrew](https://brew.sh/))
- **Linux**: `sudo apt-get install git` (Ubuntu/Debian) or `sudo yum install git` (Fedora)

### Step 2: Clone the Repository
Cloning means downloading a copy of this repository to your computer:

```bash
# This command creates a copy of the repository on your computer
git clone [repository-url]

# Navigate into the repository folder
cd code-academy
```

### Step 3: Create Your Personal Folder
Create a folder with your name (use lowercase and hyphens instead of spaces):

```bash
# Create your folder (replace 'your-name' with your actual name)
mkdir your-name

# Example: if your name is John Doe
mkdir john-doe
```

### Step 4: Create Lesson Folders
For each lesson, create a new folder inside your personal folder:

```bash
# Navigate to your folder
cd your-name

# Create a lesson folder
mkdir lesson-01

# Navigate into the lesson folder
cd lesson-01
```

## 📝 Submitting Your Work

### Basic Git Workflow

Here's how to save and submit your work. Think of this as a 3-step process:

1. **Stage** - Select which files to save
2. **Commit** - Save the files with a description
3. **Push** - Upload to the online repository

```bash
# 1. Check what files have changed
git status
# This shows you what files are new or modified

# 2. Add your files to the staging area
git add .
# The dot (.) means "add all changed files in current folder"
# You can also add specific files: git add filename.js

# 3. Commit your changes with a meaningful message
git commit -m "Add lesson 1 hello world exercise"
# The message should describe WHAT you did and WHY

# 4. Push your changes to GitHub
git push origin main
# This uploads your changes so instructors can see them
```

### 💡 Commit Message Tips

Good commit messages help everyone understand what changed:

```bash
# ✅ Good commit messages:
git commit -m "Complete lesson 1 variables exercise"
git commit -m "Fix syntax error in functions.js"
git commit -m "Add solution for array sorting challenge"

# ❌ Avoid vague messages like:
git commit -m "Fixed stuff"
git commit -m "Updates"
git commit -m "asdfasdf"
```

## 📋 Naming Conventions

Consistent naming helps keep the repository organized:

### Folder Names
- Use **lowercase letters**
- Use **hyphens** instead of spaces
- Be descriptive

```
✅ Good: john-doe, lesson-01, bonus-exercises
❌ Bad: JohnDoe, Lesson 1, my stuff
```

### File Names
- Use **lowercase letters**
- Use **hyphens** or **underscores**
- Include the file extension (.js, .py, .html, etc.)
- Be descriptive about what the file contains

```
✅ Good: hello-world.js, calculate_average.py, index.html
❌ Bad: HelloWorld.js, file1.py, untitled.txt
```

## 🛠️ Common Git Commands Reference

```bash
# Check current status
git status

# View commit history
git log --oneline

# Pull latest changes from repository
git pull origin main

# Create and switch to a new branch (advanced)
git checkout -b feature-name

# Switch between branches (advanced)
git checkout branch-name

# See what changed in your files
git diff

# Unstage files (if you added by mistake)
git reset HEAD filename

# Discard local changes to a file
git checkout -- filename
```

## ❓ Troubleshooting Common Issues

### "Permission Denied" Error
- Make sure you have the correct access rights to the repository
- Contact your instructor for access

### "Merge Conflict" Error
- This happens when two people edit the same file
- Ask your instructor for help resolving conflicts

### "Nothing to Commit" Message
- This means no files have changed
- Make sure you saved your files before committing

### Forgot to Pull Before Starting Work
```bash
# Save your current work
git stash

# Pull the latest changes
git pull origin main

# Restore your work
git stash pop
```

## 📚 Learning Resources

### Git & GitHub
- [GitHub's Git Handbook](https://guides.github.com/introduction/git-handbook/)
- [Interactive Git Tutorial](https://learngitbranching.js.org/)
- [Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf)

### Programming Basics
- [FreeCodeCamp](https://www.freecodecamp.org/)
- [Codecademy](https://www.codecademy.com/)
- [MDN Web Docs](https://developer.mozilla.org/)

## 🤝 Getting Help

If you're stuck:
1. **Check this README** - The answer might be here!
2. **Google the error message** - Someone else has probably had the same problem
3. **Ask a classmate** - Collaboration is encouraged!
4. **Ask your instructor** - That's what they're here for!

## 📌 Important Notes for Students

- **Always pull before you start working** to get the latest changes
- **Commit frequently** - It's better to have many small commits than one large one
- **Write meaningful commit messages** - Future you will thank present you
- **Don't commit sensitive information** like passwords or API keys
- **Test your code before committing** - Make sure it works!
- **Ask questions** - There are no stupid questions when learning

## 🎯 Assignment Submission Checklist

Before submitting each assignment, make sure:

- [ ] Your code is in the correct folder (your-name/lesson-XX/)
- [ ] All files are saved
- [ ] Your code runs without errors
- [ ] You've added all new files with `git add`
- [ ] You've committed with a descriptive message
- [ ] You've pushed your changes with `git push`
- [ ] You can see your changes on GitHub (web browser)

---

## 📖 Glossary for Beginners

- **Repository (Repo)**: A folder that contains all your project files and their history
- **Clone**: Download a copy of a repository to your computer
- **Commit**: A saved snapshot of your changes
- **Push**: Upload your commits to GitHub
- **Pull**: Download changes from GitHub to your computer
- **Branch**: A separate version of your code (like a parallel universe)
- **Merge**: Combine changes from different branches
- **Fork**: Create your own copy of someone else's repository
- **Pull Request (PR)**: Ask to merge your changes into the main repository

---

**Remember**: Everyone was a beginner once. Don't be afraid to make mistakes - that's how we learn! 🚀

Happy coding! 💻✨
