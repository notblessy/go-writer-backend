CREATE TABLE article_categories (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name VARCHAR(128),
  slug VARCHAR(128),
  description TEXT,

  is_active BOOLEAN DEFAULT 1,
  created_by VARCHAR(64),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE articles (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  slug VARCHAR(128),
  category_id INT,

  title VARCHAR(128),
  body TEXT,
  cover_image TEXT,
  status ENUM ('PUBLISHED', 'UNPUBLISHED','DRAFT') DEFAULT 'PUBLISHED',
  published_at timestamp,
  created_by VARCHAR(64),
  created_at timestamp,
  updated_at timestamp
);