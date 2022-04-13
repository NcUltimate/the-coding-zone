require 'pqueue'
require 'benchmark'

class QObj < Struct.new(:w, :i, :v, :n)
  def h
    w + i + n - v
  end

  def <=>(other)
    h <=> other.h
  end
end

def qsolution(num)
  qobj = QObj.new(1, 1, 2, num)
  q = PQueue.new { |a, b| b <=> a }
  q.push(qobj)

  count = 0
  while !q.empty?
    qobj = q.pop
    break if qobj.v == num

    if qobj.v * 2 <= num
      q.push(
        QObj.new(
          qobj.w + 1,
          qobj.i + 1,
          qobj.v * 2,
          num
        )
      )
    end

    if qobj.v + 1 <= num
      q.push(
        QObj.new(
          qobj.w + qobj.v,
          qobj.i + 1,
          qobj.v + 1,
          num
        )
      )
    end
  end
  qobj.i
end

def main
  40.times { |k| solution(k) }
end

def bsolution(num)
  num.to_s(2).chars.reduce(-2) { |s, k| s += (2 ** k.to_i) }
end

def main
  n = (10 ** 2 - 1)
  puts Benchmark.measure { bsolution(n) }

  bt1 = Time.now
  s1 = bsolution(n)
  bt2 = Time.now

  qt1 = Time.now
  s2 = qsolution(n)
  qt2 = Time.now

  puts "Your input: #{n}"
  puts "BSolution: #{s1} (took #{(bt2 - bt1).round(6).to_s}s)"
  puts "QSolution: #{s2} (took #{(qt2 - qt1).round(6).to_s}s)"
end

main