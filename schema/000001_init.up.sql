CREATE TABLE IF NOT EXISTS Skill (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Progress (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    passed INT NOT NULL DEFAULT 0,
    total INT NOT NULL,
    skill_id INT NOT NULL,
    CONSTRAINT fk_skill FOREIGN KEY(skill_id) REFERENCES Skill(id) ON DELETE CASCADE,
    CONSTRAINT valid_progress CHECK (passed <= total)
);

CREATE TABLE IF NOT EXISTS Books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    progress_id INT NOT NULL,
    CONSTRAINT fk_progress FOREIGN KEY(progress_id) REFERENCES Progress(id) ON DELETE CASCADE,
    UNIQUE (name, author)
);

CREATE TABLE IF NOT EXISTS Playlists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    url TEXT NOT NULL UNIQUE,
    progress_id INT NOT NULL,
    CONSTRAINT fk_progress FOREIGN KEY(progress_id) REFERENCES Progress(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Courses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    url TEXT NOT NULL UNIQUE,
    progress_id INT NOT NULL,
    CONSTRAINT fk_progress FOREIGN KEY(progress_id) REFERENCES Progress(id) ON DELETE CASCADE
);

CREATE INDEX idx_books_name ON Books (name);
CREATE INDEX idx_courses_url ON Courses (url);