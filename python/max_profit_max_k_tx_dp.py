import sys


def maxProfit(prices, totalDays, maxTx):

    dp = [[0 for i in range(totalDays)]
          for j in range(maxTx + 1)]

    for tx in range(1, maxTx + 1):
        for day in range(1, totalDays):
            # max ((previous day profit if not seeling that day), max(sell that day and include last day max profit))
            max = 0
            for Prevday in range(day):
                if prices[day] - prices[Prevday] + dp[tx-1][Prevday] > max:
                    max = prices[day] - prices[Prevday] + dp[tx-1][Prevday]

            dp[tx][day] = max if max > dp[tx][day-1] else dp[tx][day-1]

    return dp[maxTx][totalDays - 1]


if __name__ == "__main__":
    T = int(input())
    for t in range(T):
        K = int(input())
        N = int(input())
        prices = [x for x in sys.stdin.readline().strip().split()]
        print(maxProfit(prices, N, K))
'''
0 1 0 1 0 1
1 0 1 0 1 0
0 1 1 1 1 0
0 0 1 1 1 0
1 1 1 1 1 1
'''
