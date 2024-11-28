import json
import pandas as pd
import logging

#
df = pd.read_csv('prompt_engineering_dataset.csv')

df.drop(["Prompt_Type","Prompt_Length"],axis=1, inplace=True)
df = df.drop_duplicates()
df = df.sample(10 , replace=False)
data_list = [{"prompt": row['Prompt'], "response": row['Response']} for _, row in df.iterrows()]

# Export to a JSON file
with open('output.json', 'w', encoding='utf-8') as f:
    json.dump(data_list, f, ensure_ascii=False, indent=4)


