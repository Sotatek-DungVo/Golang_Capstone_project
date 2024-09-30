-- Create extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Insert mock data for users
-- password is "string"
INSERT INTO users (username, email, avatar_url, password, description, is_enabled, gender, created_at, updated_at)
VALUES
  ('john_doe', 'john@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQDN8yyDpXXfDiWnrOGw9nw5vS8W0771Ny_YA&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'I love sports', true, 'MALE', NOW(), NOW()),
  ('jane_smith', 'jane@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcShxWR5om2F4coiyMQP9n4gq4coJwFCUYMwHw&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Gamer girl', true, 'FEMALE', NOW(), NOW()),
  ('mike_johnson', 'mike@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS_Yn3d7sjy5kA9kqHmi0e24uay_wWsL6PQyg&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Competitive player', true, 'MALE', NOW(), NOW()),
  ('emily_brown', 'emily@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQcZLatv0zd77qmjcZabvDNvQmI6ubhs-iNlw&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Strategy game enthusiast', true, 'FEMALE', NOW(), NOW()),
  ('david_lee', 'david@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTbuiITbW16T50r_NVY3yTebgwTLVIBbsRFcw&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'RPG lover', true, 'MALE', NOW(), NOW()),
  ('sarah_wilson', 'sarah@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQcPVKIyxD1uC8gVJs4MUe1wor5IPwxo1t5jw&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Casual gamer', true, 'FEMALE', NOW(), NOW()),
  ('chris_taylor', 'chris@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRF1g30tPeQYkYnlhVsc5YtKScpUcsS1rDYHA&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'FPS pro', true, 'MALE', NOW(), NOW()),
  ('lisa_anderson', 'lisa@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT1Mu2hx06iLLVD3fPkOcRUhdZ-qFBHdErvKw&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'MOBA player', true, 'FEMALE', NOW(), NOW()),
  ('alex_martinez', 'alex@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQd-7CiCioR4fIHjjZHTAGHPj98HuGANkb2MA&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Indie game developer', true, 'MALE', NOW(), NOW()),
  ('olivia_garcia', 'olivia@example.com', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQoBbFQ6asLHH5s5s2-Fsx6tYg6I6DcQhA_ww&s', '$2a$10$NQM2bmxNg3sUlMFFahXrZOcsYf09B4NvXKec2LKkeLBZBa.C/4yvC', 'Retro game collector', true, 'FEMALE', NOW(), NOW());

-- Insert mock data for game categories
INSERT INTO game_categories (name, image_url, created_at, updated_at)
VALUES
  ('Naruto ultimate ninja storm', 'https://shared.akamai.steamstatic.com/store_item_assets/steam/apps/349040/capsule_616x353.jpg?t=1703080866', NOW(), NOW()),
  ('Megaman', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRd0-5Dd6X1AEY-Di2JoQPY1-X-ARM9S6nM_Q&s', NOW(), NOW()),
  ('RPG', 'https://assets1.ignimgs.com/2017/04/27/top-list-rpg-desktop-1493328959056.jpg', NOW(), NOW()),
  ('League of legends', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRiV_Y7kR7WDOEHBDXQxdbqqDnGluVJLY7S8g&s', NOW(), NOW()),
  ('Lien quan mobile', 'https://cdn.tgdd.vn/GameApp/3/219924/Screentshots/garena-lien-quan-mobile-thang-bai-tai-ky-nang-08-11-2021-8.png', NOW(), NOW()),
  ('Gunny', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSckfz3kJyYQBm9lkCv9Eq_V2ws85jsLm2jTg&s', NOW(), NOW()),
  ('PUBG PC', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQjhV2nbNgwaSxokluI2gO2nKjfA2K37B086Q&s', NOW(), NOW()),
  ('Call of duty', 'https://imgs.callofduty.com/content/dam/atvi/callofduty/cod-touchui/blackops6/meta/BO6_LP-meta_image.jpg', NOW(), NOW()),
  ('Cross fire', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSJddUdbMQQYWKSa93oU9F1p53VeesPuW_0mg&s', NOW(), NOW()),
  ('Asphalt', 'https://fdn.gsmarena.com/imgroot/news/18/08/asphalt-9-review/-728/gsmarena_002.jpg', NOW(), NOW());

-- Insert mock data for required skills
INSERT INTO required_skills (name, created_at, updated_at)
VALUES
  ('Aiming', NOW(), NOW()),
  ('Strategy', NOW(), NOW()),
  ('Teamwork', NOW(), NOW()),
  ('Map Awareness', NOW(), NOW()),
  ('Resource Management', NOW(), NOW()),
  ('Quick Reflexes', NOW(), NOW()),
  ('Communication', NOW(), NOW()),
  ('Problem Solving', NOW(), NOW()),
  ('Multitasking', NOW(), NOW()),
  ('Leadership', NOW(), NOW());

-- Insert mock data for games
INSERT INTO games (name, max_member, start_time, end_time, game_owner_id, game_category_id, created_at, updated_at)
VALUES
  ('Overwatch Tournament', 12, NOW() + INTERVAL '1 day', NOW() + INTERVAL '2 days', 1, 1, NOW(), NOW()),
  ('League of Legends Match', 10, NOW() + INTERVAL '2 days', NOW() + INTERVAL '3 days', 2, 2, NOW(), NOW()),
  ('World of Warcraft Raid', 25, NOW() + INTERVAL '3 days', NOW() + INTERVAL '4 days', 3, 3, NOW(), NOW()),
  ('Starcraft II 1v1', 2, NOW() + INTERVAL '4 days', NOW() + INTERVAL '5 days', 4, 4, NOW(), NOW()),
  ('FIFA Tournament', 16, NOW() + INTERVAL '5 days', NOW() + INTERVAL '6 days', 5, 5, NOW(), NOW()),
  ('Portal 2 Co-op', 2, NOW() + INTERVAL '6 days', NOW() + INTERVAL '7 days', 6, 6, NOW(), NOW()),
  ('Minecraft Adventure', 8, NOW() + INTERVAL '7 days', NOW() + INTERVAL '8 days', 7, 7, NOW(), NOW()),
  ('The Sims 4 Challenge', 1, NOW() + INTERVAL '8 days', NOW() + INTERVAL '9 days', 8, 8, NOW(), NOW()),
  ('Mario Kart Tournament', 8, NOW() + INTERVAL '9 days', NOW() + INTERVAL '10 days', 9, 9, NOW(), NOW()),
  ('Street Fighter V Tournament', 32, NOW() + INTERVAL '10 days', NOW() + INTERVAL '11 days', 10, 10, NOW(), NOW());

-- Insert mock data for game_required_skills
INSERT INTO game_required_skills (game_id, required_skill_id)
VALUES
  (1, 1), (1, 3), (1, 6),
  (2, 2), (2, 3), (2, 4),
  (3, 2), (3, 3), (3, 5),
  (4, 2), (4, 5), (4, 8),
  (5, 3), (5, 6), (5, 9),
  (6, 2), (6, 8), (6, 9),
  (7, 2), (7, 5), (7, 8),
  (8, 5), (8, 8), (8, 9),
  (9, 3), (9, 6), (9, 9),
  (10, 1), (10, 6), (10, 9);

 Insert mock data for game requests
INSERT INTO game_requests (user_id, game_id, status, created_at, updated_at)
VALUES
  (2, 1, 'pending', NOW(), NOW()),
  (3, 1, 'accepted', NOW(), NOW()),
  (4, 2, 'pending', NOW(), NOW()),
  (5, 2, 'rejected', NOW(), NOW()),
  (6, 3, 'accepted', NOW(), NOW()),
  (7, 3, 'pending', NOW(), NOW()),
  (8, 4, 'accepted', NOW(), NOW()),
  (9, 5, 'pending', NOW(), NOW()),
  (10, 6, 'accepted', NOW(), NOW()),
  (1, 7, 'pending', NOW(), NOW());