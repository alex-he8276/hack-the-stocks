from flask import Flask, request
import yfinance as yf
import pandas as pd
from datetime import datetime, date, timedelta

# Import Buf
import sys
sys.path.insert(1, '../ProtoBuf/gen/python/happystock/')
import happystockapi_pb2 as happystock

app = Flask(__name__)

def fetchStockPrices(ticker):

    stockPriceList = happystock.listStockPrice()

    tickerInfo = yf.Ticker(ticker)

    todayDate = date.today() + timedelta(days=1)
    weekAgoDate = date.today() - timedelta(days=6)
    stockHistory = tickerInfo.history(start=weekAgoDate.strftime("%Y-%m-%d"), end=todayDate.strftime("%Y-%m-%d"), interval="1d")

    for time, stockData in stockHistory.iterrows():
        stockItem = stockPriceList.priceList.add()
        stockItem.name = ticker 
        stockItem.price = round(stockData['Close'], 2)
    return stockPriceList

@app.route("/stock_price", methods=['GET'])
def returnStockPrices():
    stockPriceList = {}
    if 'ticker' in request.args: 
        ticker = request.args['ticker']
        stockPriceList = fetchStockPrices(ticker)
    return stockPriceList.SerializeToString()
