# Guide to use code generator 

model: City 
type: api --> api / cronjob / event
fields:
  - name: ID
    type: int
    binding: required
    constraint: primaryKey;autoIncrement
  - name: province
    type: string
    binding: required
    constraint: notNull
  - name: code
    type: int
    binding: none
    constraint: notNull
  - name: countryID
    type: int
    binding: none
    constraint: notNull
  - name: country
    type: Country
    binding: none
    constraint: notNull
routers:
  - method: get
    call: GetList
    path: /city 
    instruction: Get list the city
  - method: post
    call: Insert
    path: /city
    instruction: Function to add new city