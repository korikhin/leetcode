import polars as pl

def find_employees(df: pl.DataFrame) -> pl.DataFrame:
    return (
        df.join(
            df.select(['id', 'salary']),
            left_on='managerId',
            right_on='id',
            suffix='_manager'
        )
        .filter(pl.col('salary') > pl.col('salary_manager'))
        .select('name')
        .rename({'name': 'Employee'})
    )
