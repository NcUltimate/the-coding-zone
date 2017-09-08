require "active_support/all"
require 'set'
require 'benchmark'
class Grid
  attr_accessor :rows, :cols, :matrix, :words
  def initialize(file)
    i = 0
    j = 0
    k = 0

    File.open(file, "r") do |f|
      f.each_line do |line|
        if i == 0
          @rows, @cols = line.split(/ /)
          self.matrix = Array.new(@rows.to_i){ Array.new(@cols.to_i) }
          self.words = Array.new
        elsif i < @rows.to_i + 1
          self.matrix[i-1] = line.delete("\n").split(//)
        else
          self.words << line.delete("\n")
        end
        i += 1
      end
    end
  end
end

class Wordsearch
  class << self
    def search(grid)
      s = Set.new
      grid.words.each do |word|
        s.merge(search_horizontal(word, grid))
        s.merge(search_vertical(word, grid))
        s.merge(search_diagonal(word, grid))
        s.merge(search_row_diagonal_rev(word, grid))
        s.merge(search_col_diagonal_rev(word, grid))
      end
      print_grid(grid.matrix, s)
    end

    def print_grid(matrix, s)
      j = 0
      matrix.each do |r|
        r.each do |c|
          print s.member?(j) ? "\e[31m#{c}\e[0m" : c
          print " "
          j += 1
        end
        puts
      end
    end

    def found_match?(row, i, word)
      row[i...(word.length+i)].join == word || row[i...(word.length+i)].join == word.reverse
    end

    def search_horizontal(word, grid)
      s = Set.new
      grid.matrix.each_with_index do |row, idx|
        match1 = (row.join) =~ /#{word}/
        match2 = (row.join) =~ /#{word.reverse}/
        if match1 || match2
          start_char_index = idx * row.length + (match1 || match2)
          last_char_index = idx * row.length + word.length + (match1 || match2)
          s.merge(((start_char_index)...(last_char_index)).to_set)
        end
      end
      s
    end

    def search_vertical(word, grid)
      s = Set.new
      grid.matrix.transpose.each_with_index do |row, idx|
        match1 = (row.join) =~ /#{word}/
        match2 = (row.join) =~ /#{word.reverse}/
        if match1 || match2
          i = match1 || match2
          s.merge((i...(word.length + i)).map { |n| n * grid.cols.to_i + idx }.to_set)
        end
      end
      s
    end

    def search_diagonal(word, grid)
      k = 0
      s = Set.new
      row_set = Set.new
      start_num = nil
      row_set = search_row_diagonal(word, grid)
      s.merge(row_set) unless row_set.empty?

      col_set = Set.new
      col_set = search_col_diagonal(word, grid)
      s.merge(col_set) unless col_set.empty?

      return s
    end

    def search_row_diagonal_rev(word, grid)
      k = 0
      s = Set.new
      l = 0
      str = ""
      (k...grid.rows.to_i).reverse_each do |i|
        l = 0
        (0..i).reverse_each do |j|
          str += grid.matrix[j][l]
          l += 1
        end

        match1 = str =~ /#{word}/
        match2 = str =~ /#{word.reverse}/
        if match1 || match2
          match = match1 || match2
          start_pos = (i - match) * grid.cols.to_i + match
          s.merge(reverse_diagonal_elements(word, start_pos, grid.cols.to_i))
        end
        str = ""
      end
      s
    end

    def search_col_diagonal_rev(word, grid)
      k = 0
      s = Set.new
      l = grid.rows.to_i - 1
      str = ""
      start_pos = -1
      (k...grid.rows.to_i).reverse_each do |i|
        l = grid.rows.to_i - i - 1
        m = grid.rows.to_i - 1
        start_pos += 1
        n = l
        (l...grid.cols.to_i).each do |j|
          str += grid.matrix[m][l]
          l += 1
          m -= 1
        end
        match1 = str =~ /#{word}/
        match2 = str =~ /#{word.reverse}/
        if match1 || match2
          match = match1 || match2
          start_num = (grid.rows.to_i - 1 - match) * grid.cols.to_i + start_pos +  match
          s.merge(reverse_diagonal_elements(word, start_num, grid.cols.to_i))
        end
        str = ""
      end
      return s
    end

    def search_col_diagonal(word, grid)
      k = 0
      s = Set.new
      l = 0
      start_pos = -1
      (k...grid.cols.to_i).each do |i|
        str = ""
        l += 1
        start_pos += 1
        str = (i...grid.cols.to_i).collect do |j|
          grid.matrix[j-k][j] if j < grid.cols.to_i && (j-k) < grid.rows.to_i
        end.join
        match1 = str =~ /#{word}/
        match2 = str =~ /#{word.reverse}/
        if match1 || match2
          match = match1 || match2
          start_num = start_pos + grid.cols.to_i * match + match
          s.merge(diagonal_elements(word, start_num, grid.cols.to_i))
        end
        k += 1
      end
      s
    end

    def search_row_diagonal(word, grid)
      k = 0
      s = Set.new
      l = -1
      (0...grid.rows.to_i).each do |i|
        str = ""
        l += 1

        str =  (i...grid.rows.to_i).collect { |j| grid.matrix[j][j-k]}.join

        match1 = str =~ /#{word}/
        match2 = str =~ /#{word.reverse}/
        if match1 || match2
          match = match1 || match2
          start_pos = l * grid.cols.to_i + match * grid.cols.to_i +  match
          s.merge(diagonal_elements(word, start_pos, grid.cols.to_i))
        end
        k += 1
      end
      s
    end

    def diagonal_elements(word, start, cols)
      arr = Array.new
      arr << start
      temp = start
      (2..(word.length)).each do |_|
        temp = temp + cols + 1
        arr << temp
      end
      arr.to_set
    end

    def reverse_diagonal_elements(word, start, cols)
      arr = Array.new
      arr << start
      temp = start
      (2..(word.length)).each do |_|
        temp = temp - cols + 1
        arr << temp
      end
      arr.to_set
    end
 end
end
