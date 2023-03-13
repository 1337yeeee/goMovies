-- Add a new user to the table
INSERT INTO users (name, email, password) VALUES ('John Doe', 'john.doe@example.com', 'password');

-- Add another user to the table
INSERT INTO users (name, email, password) VALUES ('Jane Smith', 'jane.smith@example.com', 'password123');

-- Add a new director to the table
INSERT INTO directors (id, name, img, description) VALUES (1, 'Steven Spielberg', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/Steven_Spielberg_by_Gage_Skidmore.jpg/220px-Steven_Spielberg_by_Gage_Skidmore.jpg', 'American film director, producer, and screenwriter');

-- Add another director to the table
INSERT INTO directors (id, name, img, description) VALUES (2, 'Christopher Nolan', 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Christopher_Nolan_Cannes_2018.jpg/220px-Christopher_Nolan_Cannes_2018.jpg', 'British film director, producer, and screenwriter');

INSERT INTO directors (id, name, img, description) VALUES (3, 'Quentin Tarantino', 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/Quentin_Tarantino_by_Gage_Skidmore.jpg/220px-Quentin_Tarantino_by_Gage_Skidmore.jpg', 'Quentin Jerome Tarantino is an American filmmaker and screenwriter, known for his nonlinear storytelling and stylized violence.');

-- Add a new movie to the table
INSERT INTO movies (name, year, country, description, img, director_id) VALUES ('Jurassic Park', 1993, 'USA', 'A theme park suffers a major power breakdown that allows its cloned dinosaur exhibits to run amok', 'https://upload.wikimedia.org/wikipedia/en/e/e7/Jurassic_Park_poster.jpg', 1);

-- Add another movie to the table
INSERT INTO movies (name, year, country, description, img, director_id) VALUES ('The Dark Knight', 2008, 'USA', 'Batman confronts the Joker, a master criminal who is terrorizing Gotham City', 'https://upload.wikimedia.org/wikipedia/en/1/1c/The_Dark_Knight_%282008_film%29.jpg', 2);

-- Add a new movie to user 1's watchlist
INSERT INTO user_movie_ratings (user_id, movie_id, watched) VALUES (1, 1, FALSE);

-- Add a rating for movie 1 by user 1
UPDATE user_movie_ratings SET rating = 4 WHERE user_id = 1 AND movie_id = 1;

-- Add another movie to user 2's watchlist
INSERT INTO user_movie_ratings (user_id, movie_id, watched) VALUES (2, 2, FALSE);

-- Mark movie 2 as watched by user 2 and add a rating
UPDATE user_movie_ratings SET watched = TRUE, rating = 5 WHERE user_id = 2 AND movie_id = 2;

-- Add another movie to the table
INSERT INTO movies (name, year, country, description, img, director_id) VALUES ('Forrest Gump', 1994, 'USA', 'A simple man with a low IQ embarks on an extraordinary journey through life, meeting historical figures along the way', 'https://upload.wikimedia.org/wikipedia/en/thumb/6/67/Forrest_Gump_poster.jpg/220px-Forrest_Gump_poster.jpg', 1);

-- Add another movie to the table
INSERT INTO movies (name, year, country, description, img, director_id) VALUES ('Inception', 2010, 'USA', 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O', 'https://upload.wikimedia.org/wikipedia/en/2/2e/Inception_%282010%29_theatrical_poster.jpg', 2);

-- Migration 1: Add the movie "Pulp Fiction"
INSERT INTO movies (name, year, country, description, img, director_id)
VALUES ('Pulp Fiction', 1994, 'United States', 'The lives of two hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.', 'https://upload.wikimedia.org/wikipedia/en/3/3b/Pulp_Fiction_%281994%29_poster.jpg', 3);

-- Migration 2: Add the movie "Kill Bill: Vol. 1"
INSERT INTO movies (name, year, country, description, img, director_id)
VALUES ('Kill Bill: Vol. 1', 2003, 'United States', 'After awakening from a four-year coma, a former assassin wreaks vengeance on the team of assassins who betrayed her.', 'https://upload.wikimedia.org/wikipedia/en/thumb/2/2c/Kill_Bill_Volume_1.png/220px-Kill_Bill_Volume_1.png', 3);

-- Migration 3: Add the movie "Kill Bill: Vol. 2"
INSERT INTO movies (name, year, country, description, img, director_id)
VALUES ('Kill Bill: Vol. 2', 2004, 'United States', 'The Bride continues her quest of vengeance against her former boss and lover Bill, the reclusive bouncer Budd, and the treacherous, one-eyed Elle.', 'https://upload.wikimedia.org/wikipedia/en/thumb/c/c4/Kill_Bill_Volume_2.png/220px-Kill_Bill_Volume_2.png', 3);

INSERT INTO movies (name, year, country, description, img, director_id)
VALUES ('Reservoir Dogs', 1992, 'United States', 'After a simple jewelry heist goes terribly wrong, the surviving criminals begin to suspect that one of them is a police informant.', 'https://upload.wikimedia.org/wikipedia/en/f/f6/Reservoir_dogs_ver1.jpg', 3);

INSERT INTO movies (name, year, country, description, img, director_id)
VALUES ('Inglourious Basterds', 2009, 'United States', "In Nazi-occupied France during World War II, a plan to assassinate Nazi leaders by a group of Jewish U.S. soldiers coincides with a theatre owner's vengeful plans for the same.", 'https://upload.wikimedia.org/wikipedia/en/c/c3/Inglourious_Basterds_poster.jpg', 3);
