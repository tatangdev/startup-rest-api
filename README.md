# Welcome to StackEdit!

Hi! I'm your first Markdown file in **StackEdit**. If you want to learn about StackEdit, you can read me. If you want to play with Markdown, you can edit me. Once you have finished with me, you can create new files by opening the **file explorer** on the left corner of the navigation bar.

## Entity

The file explorer is accessible using the button in left corner of the navigation bar. You can create a new file by clicking the **New file** button in the file explorer. You can also create folders by clicking the **New folder** button.

### Users

| Column Name       | Data Type |
| ----------------- | --------- |
| id                | int       |
| name              | varchar   |
| occupation        | varchar   |
| email             | varchar   |
| password_hash     | varchar   |
| avatar_file_name  | varchar   |
| role              | varchar   |
| token             | varchar   |
| created_at        | datetime  |
| updated_at        | datetime  |

### Campaigns

| Column Name       | Data Type |
| ----------------- | --------- |
| id                | int       |
| user_id           | int       |
| name              | varchar   |
| short_description | varchar   |
| description       | text      |
| goal_amount       | int       |
| current_amount    | int       |
| backer_count      | int       |
| perks             | text      |
| slug              | varchar   |
| created_at        | datetime  |
| updated_at        | datetime  |

### Campaign Images

| Column Name       | Data Type |
| ----------------- | --------- |
| id                | int       |
| campaign_id       | int       |
| file_name         | varchar   |
| is_primary        | boolean   |
| created_at        | datetime  |
| updated_at        | datetime  |

### Transactions

| Column Name       | Data Type |
| ----------------- | --------- |
| id                | int       |
| campaign_id       | int       |
| user_id           | int       |
| amount            | int   |
| status            | varchar   |
| code              | varchar   |
| created_at        | datetime  |
| updated_at        | datetime  |