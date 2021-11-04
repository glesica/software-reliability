class Pair<T> {
  final T left;
  final T right;

  Pair(this.left, this.right);
}

num addPair<T extends num>(Pair<T> p) {
  return p.left + p.right;
}

void foo(int x) {
  
}
