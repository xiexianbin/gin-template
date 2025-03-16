# pagination

base on [gorm](https://gorm.io/)'s pagination plugin, feature:

- **Paging**
  - `page`: current page (default 1)
  - `page_size`: Records per page(default 10)

- **Order**
  - `order_by`: Sort fields and directions(e.g. `created_at desc, name`)

- **Filter Condition**
  - format: `<field>_<operator>=<value>`
    - e.g.: `age_gt=20` (age > 20), `name_like=John%`(LIKE 'John%')
  - support operator
    - `eq` (default)
    - `ne`
    - `gt`
    - `gte`
    - `lt`
    - `lte`
    - `like`
    - `in`

## demo

```
# URL
GET /users?page=2&page_size=20&order_by=created_at desc,name&age_gt=25&name_like=John%

# SQL
SELECT * FROM users WHERE age > 25 AND name LIKE 'John%' ORDER BY created_at DESC, name ASC LIMIT 20 OFFSET 20;
```
