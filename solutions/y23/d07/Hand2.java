package y23.d07;

import java.util.regex.Pattern;

public class Hand2 implements Comparable<Hand2> {
  private final String cards;
  private final int bid;
  private final Value value;
  private final char[] possibleCards = "J23456789TQKA".toCharArray();
  public Hand2(String cards, int bid) {
    this.cards = cards;
    this.bid = bid;

    Value best = Value.HIGH_CARD;
    for (char c : possibleCards) {
      Value withJoker = getValueFromCardsString(cards.replace('J', c));
      if (withJoker.ordinal() > best.ordinal()) {
        best = withJoker;
      }
    }
    this.value = best;
  }

  private Value getValueFromCardsString(String cards) {
    String sortedCards = cards.chars().sorted().mapToObj(c -> String.valueOf((char) c)).reduce("", String::concat);

    Pattern fiveOfAKind = Pattern.compile("(.)\\1{4}");
    if (fiveOfAKind.matcher(sortedCards).find()) {
      return Value.FIVE_OF_A_KIND;
    }

    Pattern fourOfAKind = Pattern.compile("(.)\\1{3}");
    if (fourOfAKind.matcher(sortedCards).find()) {
      return Value.FOUR_OF_A_KIND;
    }

    Pattern fullHouse = Pattern.compile("(.)\\1{1}(.)\\2{2}|(.)\\3{2}(.)\\4{1}");
    if (fullHouse.matcher(sortedCards).find()) {
      return Value.FULL_HOUSE;
    }

    Pattern threeOfAKind = Pattern.compile("(.)\\1{2}");
    if (threeOfAKind.matcher(sortedCards).find()) {
      return Value.THREE_OF_A_KIND;
    }

    Pattern twoPair = Pattern.compile("(.)\\1{1}.*(.)\\2{1}");
    if (twoPair.matcher(sortedCards).find()) {
      return Value.TWO_PAIR;
    }

    Pattern onePair = Pattern.compile("(.)\\1{1}");
    if (onePair.matcher(sortedCards).find()) {
      return Value.ONE_PAIR;
    }

    return Value.HIGH_CARD;
  }

  public int getBid() {
    return this.bid;
  }

  private String getCards() {
    return this.cards;
  }

  private Value getValue() {
    return this.value;
  }

  private int getCardValueAt(int index) {
    char card = this.getCards().charAt(index);
    if (card == 'A') {
      return 13;
    } else if (card == 'K') {
      return 12;
    } else if (card == 'Q') {
      return 11;
    } else if (card == 'J') {
      return 1;
    } else if (card == 'T') {
      return 10;
    } else {
      return Integer.parseInt(String.valueOf(card));
    }
  }

  @Override
  public int compareTo(Hand2 o) {
    if (!this.getValue().equals(o.getValue())){
      return this.getValue().ordinal() > o.getValue().ordinal() ? 1 : -1;
    }

    for (int i = 0; i < this.getCards().length(); i++) {
      if (!(this.getCards().charAt(i) == o.getCards().charAt(i))) {
        return this.getCardValueAt(i) > o.getCardValueAt(i) ? 1 : -1;
      }
    }

    return 0;
  }

  private enum Value {
    HIGH_CARD,
    ONE_PAIR,
    TWO_PAIR,
    THREE_OF_A_KIND,
    FULL_HOUSE,
    FOUR_OF_A_KIND,
    FIVE_OF_A_KIND
  }
}
