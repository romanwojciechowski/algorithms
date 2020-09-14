#include <iostream>

int binary_search_algorithm(int lista[], int element, int sizeoft)
{
    int low = 0;
    int high = sizeoft - 1;

    while (low <= high)
    {
        int mid = (low + high) / 2;
        int guess = lista[mid];

        if (guess == element)
        {
            return mid;
        }
        if (guess > element)
            high = mid - 1;
        else
            low = mid + 1;
    }
}
int main()
{
    int liczby[] = {1, 4, 8, 10, 12, 15, 17, 20, 23, 27, 30, 35};

    std::cout << binary_search_algorithm(liczby, 12, sizeof(liczby)/sizeof(int));
}
