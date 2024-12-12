import pandas as pd

def order_scores(df: pd.DataFrame) -> pd.DataFrame:
    df = df.sort_values(by=['score'], ascending=False)
    df['rank'] = df['score'].rank(method='dense', ascending=False)
    return df[['score', 'rank']]
