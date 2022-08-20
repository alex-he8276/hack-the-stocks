from flask import Flask, request
import yfinance as yf
from datetime import date, timedelta
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
        stockItem.date.FromDatetime(time) 
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
