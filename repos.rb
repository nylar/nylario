require 'json'
require 'net/http'


def normalise(name)
    normalised = name.downcase
    normalised = normalised.gsub(' ', '-')
    normalised = normalised.gsub('+', 'p')
    normalised = normalised.gsub('#', '-sharp')
end


response = Net::HTTP.get(URI('https://api.github.com/users/nylar/repos'))
json_hash = JSON.parse(response)

repos = Array.new

json_hash.each do |repo|
    repository = Hash.new

    repository['name'] = repo['name']
    repository['url'] = repo['html_url']
    repository['description'] = repo['description']
    repository['language'] = repo['language']
    repository['slug'] = normalise(repo['language'])

    repos.push(repository)
end

File.write('projects.json', JSON.dump(repos))