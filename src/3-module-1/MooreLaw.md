# Moore`s Law Definition

Moore`s Law – an observation made by Gordon Moore, according to which the amount of transistors should double about every two years. This would result in exponential increase in speed of the computers, since with the exponential increase in transistor density transistor size will shrink and smaller transistors switch faster.

# Why Moore`s Law has stopped being true

Moore's Law is no longer used due to some physical drawbacks and limitations:
- Despite the fact that smaller transistors should consume less power, density of the transistors increases exponentially as they double. This actually leads to higher power consumption.
- Higher power consumption leads to higher temperatures. Cooling can be used to solve the problem, but it only helps so much, so the processor may eventually melt, which is why it is probably the biggest drawback of the law.
- Dynamic power is proportional to the square of the voltage swing, so the larger the voltage swing is, the more power is consumed. According to Dennard Scaling, however, voltage swing should scale down with the transistor size, thus, decreasing power usage. This law, however, has its own set of problems:
    - The voltage swing cannot be deceased forever because there still has to be a threshold maintained so that the transistors could switch.
    - In any system, there is noise problems, but they become much more apparent at small voltage swings, since you cannot easily differentiate between which values are supposed to be high and which are low.
    - Leakage power also becomes more apparent as the insulators continue to scale with the transistors. As the insulators scale down, leakages occur more often.

# Conclusion

In conclusion, according to the Moore`s Law, increase in transistor density about every two years should result in exponential increase in speed, however the law is becoming irrelevant. Main reason for that power consumption increases with the increase in transistor density which leads to the rise of temperatures of the processor. The main methods of dealing with the problem contain their own limitations: coolers can only do so much and decreasing voltage swing cannot be done forever as threshold must be maintained for transistors to switch, the system becomes more error prone to the noise with smaller voltage swing and Dennard Scaling does not consider leakage.