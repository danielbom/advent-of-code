#include <stdio.h>

int OPTIMIZED = 1;

long long int a, b, c, d, e, f, g, h, x;

void echo(char *label) {
  printf("%s: a=%lld b=%lld c=%lld d=%lld e=%lld f=%lld g=%lld h=%lld\n", label,
         a, b, c, d, e, f, g, h);
}

int is_prime(int x) {
  if (x <= 3)
    return 1;
  if (x % 2 == 0)
    return 0;
  if (x % 3 == 0)
    return 0;

  int i = 5;
  int w = 2;
  while (i * i <= x) {
    if (x % i == 0)
      return 0;
    i += w;
    w = 6 - w;
  }
  return 1;
}

int main() {
  a = b = c = d = e = f = g = h = 0;
  a = 1;

  b = 57;
  c = b;
  if (a) {
    b = (b * 100) + 100000;
    c = b;
    c = c + 17000;
  }

  echo("START");
  while (1) {
    f = 1;
    d = 2;

    if (OPTIMIZED) {
      f = is_prime(b);
    } else {
      do {
        e = 2;
        do {
          g = d * e - b;

          if (!g) {
            echo("SET F");
            f = 0;
          }

          e = e + 1;
          g = e - b;
        } while (g);
        d = d + 1;
        g = d - b;
      } while (g);
    }

    if (!f) {
      h = h + 1;
    }

    g = b - c;
    if (!g) {
      break;
    }
    b = b + 17;
    echo("NEXT");
  }
  echo("END");

  return 0;
}
