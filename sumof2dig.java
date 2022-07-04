import java.util.Scanner;

public class sumof2dig {
    public static void main (String[]args)
    {
        Scanner tr = new Scanner (System.in);
        System.out.println("Enter Your First Number");
        float a = tr.nextInt();
        System.out.println("Enter Your Next Number ");
        float b = tr.nextInt();
        double c=(a+b);
        System.out.println("Sum Is "+c);
    }
}
