from flask import Flask, request
import yfinance as yf
import pandas as pd
from datetime import date, timedelta
import happystockapi_pb2 as happystock


app = Flask(__name__)

def fetchStockPrices(ticker):

    stockPriceList = happystock.listStockPrice()

    tickerInfo = yf.Ticker(ticker)

    todayDate = date.today() + timedelta(days=1)
    weekAgoDate = date.today() - timedelta(days=8)
    stockHistory = tickerInfo.history(start=weekAgoDate.strftime("%Y-%m-%d"), end=todayDate.strftime("%Y-%m-%d"), interval="1d")

 
    for i in range(7):
        day = date.today() - timedelta(days=6-i) 

        stockItem = stockPriceList.priceList.add()
        stockItem.name = ticker   
        stockItem.date.FromDatetime(pd.to_datetime(day))

        while day.strftime("%Y-%m-%d") not in stockHistory.index:
            day = day - timedelta(days=1)

        dayStockData = stockHistory.loc[day.strftime("%Y-%m-%d")]
        stockItem.price = round(dayStockData['Close'], 2)

    return stockPriceList

@app.route("/stock_price", methods=['GET'])
def returnStockPrices():
    stockPriceList = {}
    if 'ticker' in request.args: 
        ticker = request.args['ticker']
        stockPriceList = fetchStockPrices(ticker)
    headers = {'Access-Control-Allow-Origin': '*'}
    data = stockPriceList.SerializeToString()
    return data, headers
