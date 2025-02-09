import polars as pl

def capitalize_content(user_content: pl.DataFrame) -> pl.DataFrame:
    return user_content.with_columns(
        converted_text = pl.col('content_text').str.to_titlecase()
    ).rename({'content_text': 'original_text'})
