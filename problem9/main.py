def rectArea(c1, c2):
    v = (abs(int(c1[0]) - int(c2[0])) + 1) * (abs(int(c1[1]) - int(c2[1])) + 1)
    print("c1", c1, "c2", c2, "area", v)
    return v

def main():
    with open('input.txt') as f:
            coords = [tuple(line.rstrip().split(',')) for line in f]
    # print(coords)
    maxvalue = 0
    for i in range(len(coords)):
         for j in range(len(coords)):
              if j == i:
                   continue
              maxvalue = max(maxvalue, rectArea(coords[i], coords[j]))
    
    print(maxvalue)


if __name__=="__main__":
    main()
