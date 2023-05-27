import json
import csv

print("Test")
csv_file_path = 'Pincode.csv'
csv_data =[]

with open(csv_file_path,'r') as csv_file:
    csv_reader = csv.DictReader(csv_file)
    csv_data = list(csv_reader)

json_file_path = 'pin.json'
json_data ={}

with open(json_file_path ,'r') as json_file:
    json_data =json.load(json_file)


for item in csv_data:
    item_id = item['Pincode']
    item_district = item['District']
    item_state = item['StateName']

    if item_id in json_data:
        district = item_district.split()
        lower_words =[word.lower() for word in district]
        capital_words =[word.capitalize() for word in lower_words]
        res = " ".join(capital_words)
            
        state = item_state.split()
        lower_s_word = [word.lower() for word in state]
        cap_word = [word.capitalize() for word in lower_s_word]
        
        json_data[item_id]['city'] = res

        
        json_data[item_id]['state'] = item_state

with open(json_file_path,'w') as json_file:
    json.dump(json_data,json_file,indent=2)



