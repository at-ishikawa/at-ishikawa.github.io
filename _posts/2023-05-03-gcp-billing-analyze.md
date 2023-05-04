---
title:
date: 2023-05-23
tags:
  - gcp
  - billing
---

There are a few documents to manage billing data in BigQuery
- [Attribution of committed use discount fees and credits](https://cloud.google.com/docs/cuds-attribution)
- [How to export to BigQuery](https://cloud.google.com/billing/docs/how-to/export-data-bigquery)
- [Structure of the standard data](https://cloud.google.com/billing/docs/how-to/export-data-bigquery-tables/standard-usage)

# Queries

**Sum of all costs**
From [the official document](https://cloud.google.com/billing/docs/how-to/bq-examples#sum-costs-per-invoice):

```sql
SELECT
  invoice.month,
  SUM(cost)
    + SUM(IFNULL((SELECT SUM(c.amount)
                  FROM UNNEST(credits) c), 0))
    AS total,
  (SUM(CAST(cost AS NUMERIC))
    + SUM(IFNULL((SELECT SUM(CAST(c.amount AS NUMERIC))
                  FROM UNNEST(credits) AS c), 0)))
    AS total_exact
FROM `project.dataset.gcp_billing_export_v1_XXXXXX_XXXXXX_XXXXXX`
GROUP BY 1
ORDER BY 1 ASC
;
```

**How to calculate commitment fees**
From [the official document](https://cloud.google.com/billing/docs/how-to/bq-examples#cud-fees):

```sql
SELECT
    invoice.month AS invoice_month,
    SUM(cost) as commitment_fees
FROM `project.dataset.gcp_billing_export_v1_XXXXXX_XXXXXX_XXXXXX`
WHERE LOWER(sku.description) LIKE "commitment%"
GROUP BY 1
```

**How to calculate commitment credits**

From [the official document](https://cloud.google.com/billing/docs/how-to/bq-examples#cud-credits):

```sql
SELECT
    invoice.month AS invoice_month,
    SUM(credits.amount) as CUD_credits
FROM `project.dataset.gcp_billing_export_v1_XXXXXX_XXXXXX_XXXXXX`
LEFT JOIN UNNEST(credits) AS credits
WHERE credits.type = "COMMITTED_USAGE_DISCOUNT"
GROUP BY 1
```
