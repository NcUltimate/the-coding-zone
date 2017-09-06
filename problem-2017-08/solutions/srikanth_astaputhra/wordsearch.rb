require "active_support/all"
require 'set'
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
      count = s.size
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
        (0..row.length).each do |i|
          start_char_index = idx * row.length + i
          last_char_index = idx * row.length + word.length + i
          s.merge(((start_char_index)...(last_char_index)).to_set) if found_match?(row, i, word)
        end
      end
      s
    end

    def search_vertical(word, grid)
      s = Set.new
      grid.matrix.transpose.each_with_index do |row, idx|
        (0..row.length).each do |i|
          s.merge((i...(word.length + i)).map { |n| n * grid.cols.to_i + idx }.to_set) if found_match?(row, i, word)
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
        start_num = search_reverse_diagonal_line(word, str, i * grid.cols.to_i, grid.cols.to_i)
        s.merge(reverse_diagonal_elements(word, start_num, grid.cols.to_i)) unless start_num.nil?
        str = ""
      end
      s
    end

    def search_col_diagonal_rev(word, grid)
      k = 0
      s = Set.new
      l = grid.rows.to_i - 1
      str = ""
      (k...grid.rows.to_i).reverse_each do |i|
        l = grid.rows.to_i - i - 1
        m = grid.rows.to_i - 1
        n = l
        (l...grid.cols.to_i).each do |j|
          str += grid.matrix[m][l]
          l += 1
          m -= 1

        end
        start_num = search_reverse_diagonal_line(word, str, (grid.rows.to_i - i - 1) * (grid.cols.to_i) + n, grid.cols.to_i)
        s.merge(reverse_diagonal_elements(word, start_num, grid.cols.to_i)) unless start_num.nil?
        str = ""
      end
      return s
    end

    def search_col_diagonal(word, grid)
      k = 0
      s = Set.new
      l = 0
      (k...grid.cols.to_i).each do |i|
        str = ""
        l += 1
        (i...grid.cols.to_i).each do |j|
          if j < grid.cols.to_i && (j-k) < grid.rows.to_i
            str += grid.matrix[j-k][j]
            start_num = search_diagonal_line(word, str, i, grid.cols.to_i)
            return diagonal_elements(word, start_num, grid.cols.to_i) unless start_num.nil?
          end
        end
        k += 1
      end
      s
    end

    def search_row_diagonal(word, grid)
      k = 0
      s = Set.new
      l = 0
      (0...grid.rows.to_i).each do |i|
        str = ""
        l += 1
        (i...grid.rows.to_i).each do |j|
          str += grid.matrix[j][j-k]
        end
        start_num = search_diagonal_line(word, str, i, grid.cols.to_i)
        return diagonal_elements(word, start_num * grid.cols.to_i , grid.cols.to_i) unless start_num.nil?
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

    def search_reverse_diagonal_line(word, line, start_num, offset)
      (0...line.length).each do |i|
        return start_num if found_match?(line.split(//), i, word)
        start_num = start_num - offset + 1
      end
      return nil
    end

    def search_diagonal_line(word, line, start_num, offset)
      (0...line.length).each do |i|
        return start_num if found_match?(line.split(//), i, word)
        start_num = start_num + offset + 1
      end
      return nil
    end
 end
end
