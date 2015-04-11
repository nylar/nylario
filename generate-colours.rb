require 'json'
require 'net/http'


def normalise(name)
    normalised = name.downcase
    normalised = normalised.gsub(' ', '-')
    normalised = normalised.gsub('+', 'p')
    normalised = normalised.gsub('#', '-sharp')
end


response = Net::HTTP.get(URI('https://raw.githubusercontent.com/doda/github-language-colors/master/colors.json'))
json_hash = JSON.parse(response)

css = ""

json_hash.each do |language, colour|
    normalised_language = normalise(language)
    css << "span.#{normalised_language} { color: #{colour} }\n"
end

File.write('public/css/colours.css', css)