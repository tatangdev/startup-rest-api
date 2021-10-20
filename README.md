# Welcome to StackEdit!

Hi! I'm your first Markdown file in **StackEdit**. If you want to learn about StackEdit, you can read me. If you want to play with Markdown, you can edit me. Once you have finished with me, you can create new files by opening the **file explorer** on the left corner of the navigation bar.

## Entity

The file explorer is accessible using the button in left corner of the navigation bar. You can create a new file by clicking the **New file** button in the file explorer. You can also create folders by clicking the **New folder** button.

### Users

| Column Name | Data Type | Default |
| --- | ----------- | --- |
| Header | Title | null |
| Paragraph | Text | null |

### Campaigns

| Column Name | Data Type | Default |
| --- | ----------- | --- |
| Header | Title | null |
| Paragraph | Text | null |

### Campaign Images

| Column Name | Data Type | Default |
| --- | ----------- | --- |
| Header | Title | null |
| Paragraph | Text | null |

### Transactions

| Column Name | Data Type | Default |
| --- | ----------- | --- |
| Header | Title | null |
| Paragraph | Text | null |

## UML diagrams

You can render UML diagrams using [Mermaid](https://mermaidjs.github.io/). For example, this will produce a sequence diagram:

```mermaid
sequenceDiagram
Alice ->> Bob: Hello Bob, how are you?
Bob-->>John: How about you John?
Bob--x Alice: I am good thanks!
Bob-x John: I am good thanks!
Note right of John: Bob thinks a long<br/>long time, so long<br/>that the text does<br/>not fit on a row.

Bob-->Alice: Checking with John...
Alice->John: Yes... John, how are you?
```

And this will produce a flow chart:

```mermaid
graph LR
A[Square Rect] -- Link text --> B((Circle))
A --> C(Round Rect)
B --> D{Rhombus}
C --> D
```