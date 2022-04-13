int main() {
  int n = 9999999;
  int c = 0;
  while(n != 1) {
    c += n & 1 ? 2 : 1;
    n = n >> 1;
  }
  return c;
}