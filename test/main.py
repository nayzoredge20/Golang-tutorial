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

    if item_id in json_data:
        json_data[item_id]['city'] = item_district.capitalize()
        json_data[item_id]['state'] = item['StateName'].capitalize()

with open(json_file_path,'w') as json_file:
    json.dump(json_data,json_file,indent=2)



