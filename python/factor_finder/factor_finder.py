class FactorFinder(object):
    def greatest_common_factor(self, numbers):
        return max(self.get_common_factors(numbers))

    def get_common_factors(self, numbers):
        all_factors = self.__get_all_factors(numbers)
        return set.intersection(*all_factors)

    def __get_all_factors(self, numbers):
        all_factors = []
        for num in numbers:
            all_factors.append(self.get_factors(num))
        return all_factors

    def get_factors(self, number):
        factors = set()
        for possible_factor in range(1, abs(number) + 1):
            self.__add_factors(factors, possible_factor, number)

        return factors

    def __add_factors(self, factors, possible_factor, number):
        if possible_factor in factors:
            return

        div_result, mod_result = divmod(number, possible_factor)
        if mod_result == 0:
            factors.update([possible_factor, div_result])
