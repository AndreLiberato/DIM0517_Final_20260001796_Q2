import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class CalcTest {

    @Test
    void testAdd() {
        assertEquals(7, Calc.add(3, 4));
    }

    @Test
    void testSub4_3() {
        assertEquals(1, Calc.sub(4, 3));
    }

    @Test
    void testSub3_4() {
        assertEquals(-1, Calc.sub(3, 4));
    }
}
