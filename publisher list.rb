require 'rubygems' if RUBY_VERSION < '1.9'
require 'rest_client'

headers = {
  :authorization => 'Bearer TOKEN'
}

response = RestClient.get 'https://partner.opened.com/1/publishers.json', headers
puts response
