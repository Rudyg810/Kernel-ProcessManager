solid principles

1. Single Responsibility Principle
   - A class/module should have only one reason to change
   - Each class should focus on doing one specific thing well
   - Helps maintain cleaner, more focused code
   - Makes code easier to understand and modify
   - Reduces coupling between different parts of the system

2. Open/Closed Principle
   - Software entities should be open for extension but closed for modification
   - New functionality should be added via extension not modification
   - Achieved through interfaces and abstract classes
   - Promotes code reuse and maintainability
   - Reduces risk of breaking existing code

3. Liskov Substitution Principle
   - Objects of a superclass should be replaceable with objects of subclasses
   - Subclasses must fulfill contracts of base classes
   - Ensures behavioral compatibility in inheritance hierarchies
   - Prevents unexpected bugs when using polymorphism
   - Makes code more reliable and predictable

4. Interface Segregation Principle
   - Clients should not depend on interfaces they don't use
   - Keep interfaces small and focused on specific functionality
   - Better to have many specific interfaces than one general interface
   - Reduces coupling between components
   - Makes system more flexible and maintainable

5. Dependency Inversion Principle
   - High-level modules should not depend on low-level modules
   - Both should depend on abstractions
   - Abstractions should not depend on details
   - Details should depend on abstractions
   - Enables better decoupling and flexibility
