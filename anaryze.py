import pandas as pd
import matplotlib.pyplot as plt

import matplotlib.dates as mdates

df = pd.read_json("bWOUe_X9Q8I.json")

df2 = df.groupby(pd.Grouper(key='timestamp', freq='1min')).count().reset_index()

df2 = df2[['message', 'timestamp']]

df2['message'].plot.bar()
plt.hlines([df2['message'].mean()], 0, len(df2), "blue", linestyles='dashed')
plt.show()
