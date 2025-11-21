### API Calls


https://api.deadlock-api.com/v1/analytics/hero-stats?min_unix_timestamp=1761350091&max_unix_timestamp=1763769599&min_average_badge=91&max_average_badge=116&min_hero_matches=0&min_hero_matches_total=0
- divide wins by matches for win rate
- sum heroes' match numbers divide by 12, then divide each hero's number of matches by the original calculation for pick rate

https://assets.deadlock-api.com/v2/heroes?only_active=true
- receive ids of active heroes