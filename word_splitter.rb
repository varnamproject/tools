
inputFile = ARGV.first

fileContents = File.read(inputFile)
words = fileContents.split " "

uniqueWords = {}
words.each do |word|
    uniqueWords[word] = ""
end


File.open("words.txt", 'w') do |file| 
    uniqueWords.each_key do |word|
        file.write(word)
        file.write("\n")
    end
end

puts "Wrote words.txt"
