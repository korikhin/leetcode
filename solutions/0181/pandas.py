import pandas as pd

def find_employees(df: pd.DataFrame) -> pd.DataFrame:
    return (
        df.merge(
            df[['id', 'salary']],
            left_on='managerId',
            right_on='id',
            suffixes=('', '_manager')
        )
        .loc[lambda x: x['salary'] > x['salary_manager']]
        .select('name')
        .rename(columns={'name': 'Employee'})
    )
