import requests

def get_profile(name):

    headers = {
        'authority': 'colonist.io',
        'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',
        'accept': '*/*',
        'sec-gpc': '1',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-mode': 'cors',
        'sec-fetch-dest': 'empty',
        'referer': 'https://colonist.io/profile/OPanda',
        'accept-language': 'en-US,en;q=0.9',
        'cookie': '_ga=GA1.2.73033516.1638882390; __qca=P0-69088706-1638882390876; _pbjs_userid_consent_data=3524755945110770; ucf_uid=7dcb0a47-3e7b-4ed5-a49a-a9c8396de079; __gads=ID=c47463054517a2ad:T=1638882412:S=ALNI_MaiLVw9cFDTpNmJvZo3JjiHsLb19w; __stripe_mid=f3a1b8ba-953d-4658-94de-ed503f2626f5fcbcb0; Indicative_6c6fed44-919c-432e-af43-659159a6fb82="%7B%22defaultUniqueID%22%3A%22ed6c2dcd-7ef3-4af6-8c46-a8316027ae20%22%7D"; _gid=GA1.2.921523848.1640529862; cto_bundle=8OudOl94aCUyQjV3TmI4YzZqY0lvVmJiWHZpRkVOUHpIJTJCc3p2dVolMkIlMkJEQVJTSHFlSWFPTjU4MHVSWWhiNTFsZlBsYUhlalljJTJGVjdNUG1JckQlMkYlMkY5RnphaVp3MFB0R0V3TVB0cmNrRFRHeXlBdm01Nk8lMkZoR1J3U21KOTNZZ3JwQnBHNmNXSE5wb2VxOWJGU2NDWEpuYWpqNmx1ZHZuVHJrS3NzUXRjV3VFUmhlYlUxWmdPdFZLSCUyRmo5UGMlMkJBNVFmdnpVVHV0Qw; cto_bundle=s1XZ9F94aCUyQjV3TmI4YzZqY0lvVmJiWHZpRk96NDlKOTglMkJ3YXVFdDZzaWxXUnZ6VnozbndBRHpQJTJCRzlWM3prOSUyRiUyRkx0YURMRXI2dnhnZlI2JTJCSmU5UnJmWnRnTUxydmRFcDFZR2ZBMHBCN2dmbDA5eUJqY3hvSUE0SCUyQkhabUM2dzNzUVFPY3QzVVZhYnJvMU1Hb3FGaGFFRndXQSUzRCUzRA; cto_bidid=7ojVSl9HaGRxRnZkOGtxbXpFaWFNY3V6T0kyWERVVDRxcXdZcjNuc0ZsTjc1Y294MXVSSTNsSVVERGpRN3lLJTJGZnltVXg3cUZCbFZNZWNSRTNLY3Z5amZxT0N5VkliYWNSb2w5ZG9sYnV1SXVYVVBZJTNE; jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIzMTg0MTg1NyIsImlhdCI6MTY0MTM0OTczMSwiZXhwIjoxNjQzOTQxNzMxLCJhdWQiOiJjb2xvbmlzdC5pbyIsImlzcyI6ImNvbG9uaXN0LmlvIn0.sNrWh4pIpxiVVKquVzuyt8FrD4qGRk_R9Cz5H7k70yM',
        'if-none-match': 'W/"1a700-xsJHYFSKcwRJJgJ5tj0y7oVIh7M"',
    }

    response = requests.get('https://colonist.io/api/profile/'+name, headers=headers)
    # print(response.json())
    return response.json()
