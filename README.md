# Data Environment Documentation

This environment provides access to business data through SQL queries and AI-powered analysis capabilities.

## Available Data Tables

### Public Accounts
Contains information about business accounts including:
- account_id
- account_name  
- billing_city
- billing_country
- billing_state
- industry

### Public Contacts
Contains contact information associated with accounts:
- account_id
- contact_id
- email
- first_name
- last_name
- phone

### Public Customers
Contains customer information:
- account_id
- customer_id
- email
- hvac_system_type
- name

### Public Opportunities 
Contains sales opportunity data:
- account_id
- amount (numeric)
- close_date (date)
- name
- opportunity_id
- stage

### Public Tasks
Contains task/activity information:
- due_date (date)
- priority
- related_to
- status
- subject
- task_id

## Key Capabilities

### SQL Queries
- Apache DataFusion style SQL queries supported
- Ability to join across tables
- Date/timestamp handling
- Numeric calculations

### AI Functions

#### Classification
Categorize text inputs into predefined categories:
```python
classifications = classify(
    instructions="Classification instructions",
    categories=["category1", "category2"],
    allow_multiple=False,
    inputs=["text1", "text2"]
)
```

#### Summarization
Generate summaries of text inputs:
```python
summaries = summarize(
    instructions="Summarization instructions",
    inputs=["text1", "text2"]
)
```

#### Structured Information Extraction
Extract structured data from text according to a schema:
```python
extracted_data = extract(
    json_schema=schema,
    instructions="Extraction instructions",
    inputs=["text1", "text2"]
)
```

### Data Artifacts
- Store and retrieve data artifacts for future reference
- Support for both tabular and text data
- Artifacts can be referenced in responses using <artifact /> tags

## Example Usage

Query opportunities closing this month:
```sql
SELECT 
    o.opportunity_id,
    o.name,
    o.amount,
    o.close_date,
    a.account_name
FROM 
    app.public_opportunities o
    JOIN app.public_accounts a ON o.account_id = a.account_id
WHERE 
    o.close_date >= CURRENT_DATE
    AND o.close_date < DATE_ADD(MONTH, 1, DATE_TRUNC('MONTH', CURRENT_DATE))
```

## Notes
- All timestamps are handled in UTC
- SQL queries should explicitly specify required columns rather than using SELECT *
- Use COUNT(1) instead of COUNT(*)

