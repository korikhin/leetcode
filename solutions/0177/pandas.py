import pandas as pd

def nth_highest_salary(df: pd.DataFrame, /, *, n: int = 1) -> pd.DataFrame:
    no_result = pd.DataFrame({f'getNthHighestSalary({n})': [None]})

    if n < 1:
        return no_result

    tops = (
        df['salary']
        .drop_duplicates()
        .nlargest(n)
    )

    if len(tops) < n:
        return no_result

    return pd.DataFrame({f'getNthHighestSalary({n})': [tops.iloc[-1]]})
